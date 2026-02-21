#!/usr/bin/env python3
"""
Generate a Cloudbeat CVE security statement in YAML format.

The script consumes issue context (title/body/CVE hints), runs govulncheck
against the current repository, and emits a structured YAML report that can be
posted back to the originating issue.
"""

from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import re
import subprocess
import sys
from collections import defaultdict
from dataclasses import dataclass
from typing import Any, Dict, Iterable, List, Tuple


CVE_PATTERN = re.compile(r"\bCVE-\d{4}-\d{4,7}\b", re.IGNORECASE)
SAFE_YAML_SCALAR = re.compile(r"^[A-Za-z0-9._/@:+-]+$")
RESERVED_YAML_WORDS = {"null", "true", "false", "yes", "no", "on", "off", "~"}


@dataclass(frozen=True)
class FindingView:
    osv_id: str
    reachability: str
    fixed_version: str
    trace_summary: str


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Generate Cloudbeat CVE security statement.")
    parser.add_argument("--issue-repository", default=os.getenv("ISSUE_REPOSITORY", ""))
    parser.add_argument("--issue-number", default=os.getenv("ISSUE_NUMBER", ""))
    parser.add_argument("--issue-title", default=os.getenv("ISSUE_TITLE", ""))
    parser.add_argument("--issue-body", default=os.getenv("ISSUE_BODY", ""))
    parser.add_argument("--cve-id", action="append", default=[])
    parser.add_argument("--trigger-mode", default=os.getenv("TRIGGER_MODE", "manual"))
    parser.add_argument("--source-repository", default=os.getenv("SOURCE_REPOSITORY", "elastic/cloudbeat"))
    parser.add_argument("--source-ref", default=os.getenv("SOURCE_REF", "main"))
    parser.add_argument("--output", default="security-statement.yaml")
    parser.add_argument("--skip-govulncheck", action="store_true", help="Do not execute govulncheck.")
    return parser.parse_args()


def normalize_cve_id(candidate: str) -> str:
    return candidate.strip().upper()


def collect_cve_ids(issue_title: str, issue_body: str, cli_cve_ids: Iterable[str]) -> List[str]:
    found = []
    for raw in cli_cve_ids:
        if raw:
            found.extend(CVE_PATTERN.findall(raw))
            if not CVE_PATTERN.search(raw):
                found.append(raw.strip().upper())

    found.extend(CVE_PATTERN.findall(issue_title or ""))
    found.extend(CVE_PATTERN.findall(issue_body or ""))

    deduped = []
    seen = set()
    for item in found:
        normalized = normalize_cve_id(item)
        if normalized and normalized not in seen:
            seen.add(normalized)
            deduped.append(normalized)
    return deduped


def run_govulncheck() -> Tuple[int, str, str]:
    command = ["go", "run", "golang.org/x/vuln/cmd/govulncheck@latest", "-json", "./..."]
    completed = subprocess.run(command, check=False, capture_output=True, text=True)
    return completed.returncode, completed.stdout or "", completed.stderr or ""


def parse_govulncheck_stream(stdout: str) -> Tuple[Dict[str, Dict[str, Any]], List[Dict[str, Any]]]:
    osv_by_id: Dict[str, Dict[str, Any]] = {}
    findings: List[Dict[str, Any]] = []
    for raw_line in stdout.splitlines():
        line = raw_line.strip()
        if not line:
            continue
        try:
            payload = json.loads(line)
        except json.JSONDecodeError:
            continue
        if not isinstance(payload, dict):
            continue

        osv = payload.get("osv")
        if isinstance(osv, dict) and osv.get("id"):
            osv_by_id[osv["id"]] = osv

        finding = payload.get("finding")
        if isinstance(finding, dict):
            findings.append(finding)
    return osv_by_id, findings


def classify_finding(trace: List[Dict[str, Any]]) -> str:
    has_function = any(frame.get("function") for frame in trace)
    has_package = any(frame.get("package") for frame in trace)
    has_module = any(frame.get("module") for frame in trace)

    if has_function:
        return "reachable"
    if has_package:
        return "imported_not_called"
    if has_module:
        return "dependency_present"
    return "unknown"


def summarize_trace(trace: List[Dict[str, Any]]) -> str:
    pieces: List[str] = []
    for frame in trace[:5]:
        module = frame.get("module")
        package = frame.get("package")
        function = frame.get("function")
        value = function or package or module
        if value:
            pieces.append(value)
    if not pieces:
        return "no trace details reported"
    return " -> ".join(pieces)


def combine_reachability(values: Iterable[str]) -> str:
    rank = {
        "reachable": 4,
        "imported_not_called": 3,
        "dependency_present": 2,
        "unknown": 1,
        "not_detected": 0,
    }
    best = "not_detected"
    for value in values:
        if rank.get(value, 0) > rank.get(best, 0):
            best = value
    return best


def finding_status(reachability: str, osv_matches: int, finding_count: int) -> str:
    if reachability == "reachable":
        return "affected"
    if reachability in {"imported_not_called", "dependency_present"}:
        return "potentially_affected"
    if osv_matches > 0 and finding_count == 0:
        return "not_affected"
    return "not_detected_in_go_vuln_db"


def to_yaml_scalar(value: Any) -> str:
    if value is None:
        return "null"
    if isinstance(value, bool):
        return "true" if value else "false"
    if isinstance(value, (int, float)):
        return str(value)

    text = str(value)
    if text == "":
        return '""'
    if "\n" in text:
        raise ValueError("multiline string should be handled separately")
    if SAFE_YAML_SCALAR.match(text) and text.lower() not in RESERVED_YAML_WORDS:
        return text
    return json.dumps(text)


def dump_yaml(node: Any, indent: int = 0) -> List[str]:
    pad = " " * indent

    if isinstance(node, dict):
        if not node:
            return [pad + "{}"]
        lines: List[str] = []
        for key, value in node.items():
            key_text = str(key)
            if isinstance(value, str) and "\n" in value:
                lines.append(f"{pad}{key_text}: |-")
                for block_line in value.splitlines():
                    lines.append(f"{pad}  {block_line}")
                if value.endswith("\n"):
                    lines.append(f"{pad}  ")
            elif isinstance(value, (dict, list)):
                lines.append(f"{pad}{key_text}:")
                lines.extend(dump_yaml(value, indent + 2))
            else:
                lines.append(f"{pad}{key_text}: {to_yaml_scalar(value)}")
        return lines

    if isinstance(node, list):
        if not node:
            return [pad + "[]"]
        lines = []
        for item in node:
            if isinstance(item, str) and "\n" in item:
                lines.append(pad + "- |-")
                for block_line in item.splitlines():
                    lines.append(f"{pad}  {block_line}")
            elif isinstance(item, (dict, list)):
                lines.append(pad + "-")
                lines.extend(dump_yaml(item, indent + 2))
            else:
                lines.append(f"{pad}- {to_yaml_scalar(item)}")
        return lines

    return [pad + to_yaml_scalar(node)]


def main() -> int:
    args = parse_args()

    cve_ids = collect_cve_ids(args.issue_title, args.issue_body, args.cve_id)
    now = dt.datetime.now(dt.timezone.utc).replace(microsecond=0).isoformat()

    govulncheck_stdout = ""
    govulncheck_stderr = ""
    govulncheck_exit = -1
    if not args.skip_govulncheck:
        govulncheck_exit, govulncheck_stdout, govulncheck_stderr = run_govulncheck()

    osv_by_id, findings = parse_govulncheck_stream(govulncheck_stdout)
    findings_by_osv = defaultdict(list)
    for finding in findings:
        osv_id = finding.get("osv")
        if osv_id:
            findings_by_osv[osv_id].append(finding)

    investigations: List[Dict[str, Any]] = []
    for cve_id in cve_ids:
        matching_osv = []
        for osv in osv_by_id.values():
            aliases = [alias.upper() for alias in osv.get("aliases", [])]
            if cve_id == osv.get("id", "").upper() or cve_id in aliases:
                matching_osv.append(osv)

        finding_views: List[FindingView] = []
        for osv in matching_osv:
            for finding in findings_by_osv.get(osv["id"], []):
                trace = finding.get("trace") or []
                finding_views.append(
                    FindingView(
                        osv_id=osv["id"],
                        reachability=classify_finding(trace),
                        fixed_version=finding.get("fixed_version", ""),
                        trace_summary=summarize_trace(trace),
                    )
                )

        reachability = combine_reachability([view.reachability for view in finding_views])
        status = finding_status(reachability, len(matching_osv), len(finding_views))

        summary = next((osv.get("summary", "") for osv in matching_osv if osv.get("summary")), "")
        details = next((osv.get("details", "") for osv in matching_osv if osv.get("details")), "")
        if not matching_osv:
            summary = "No matching Go advisory found for this CVE."
            details = (
                "The CVE might target a component outside of Go modules or a package "
                "that is not in the Cloudbeat dependency graph."
            )

        references = []
        for osv in matching_osv:
            for ref in osv.get("references", []):
                if isinstance(ref, dict) and ref.get("url"):
                    references.append(ref["url"])
        references = sorted(set(references))

        aliases = sorted(
            {
                alias.upper()
                for osv in matching_osv
                for alias in osv.get("aliases", [])
                if alias.upper() != cve_id
            }
        )
        fixed_versions = sorted({view.fixed_version for view in finding_views if view.fixed_version})

        investigations.append(
            {
                "cve_id": cve_id,
                "status": status,
                "reachability": reachability,
                "matching_go_advisories": [osv.get("id") for osv in matching_osv],
                "metadata": {
                    "summary": summary or "No summary returned from govulncheck advisory metadata.",
                    "details": details or "No additional details returned from govulncheck advisory metadata.",
                    "aliases": aliases,
                    "references": references,
                },
                "evidence": {
                    "findings_count": len(finding_views),
                    "fixed_versions": fixed_versions,
                    "traces": [view.trace_summary for view in finding_views[:5]],
                },
            }
        )

    statuses = {inv["status"] for inv in investigations}
    if "affected" in statuses:
        overall_assessment = "affected"
    elif "potentially_affected" in statuses:
        overall_assessment = "review_required"
    elif statuses == {"not_affected"}:
        overall_assessment = "not_affected"
    elif statuses:
        overall_assessment = "inconclusive"
    else:
        overall_assessment = "inconclusive"

    statement = {
        "schema_version": "1.0",
        "statement_type": "cloudbeat_cve_investigation",
        "generated_at": now,
        "trigger": {
            "mode": args.trigger_mode,
            "issue_repository": args.issue_repository,
            "issue_number": args.issue_number,
        },
        "analysis_context": {
            "source_repository": args.source_repository,
            "source_ref": args.source_ref,
            "issue_title": args.issue_title,
            "cloudbeat_keyword_present": "cloudbeat" in (args.issue_title + " " + args.issue_body).lower(),
            "cve_ids": cve_ids,
        },
        "tooling": {
            "govulncheck": {
                "executed": not args.skip_govulncheck,
                "exit_code": govulncheck_exit,
                "stderr": govulncheck_stderr.strip(),
            }
        },
        "overall_assessment": overall_assessment,
        "investigations": investigations,
        "recommendation": (
            "If any status is affected or potentially_affected, assign an engineer for patch planning "
            "and verify package upgrade constraints in go.mod/go.sum."
        ),
    }

    yaml_payload = "\n".join(dump_yaml(statement)) + "\n"
    with open(args.output, "w", encoding="utf-8") as handle:
        handle.write(yaml_payload)

    print(f"Wrote security statement to {args.output}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
