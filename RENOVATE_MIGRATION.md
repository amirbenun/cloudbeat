# Renovate Migration Guide

## Overview

This document outlines the migration from Dependabot to Renovate for dependency management in the cloudbeat repository.

## Why Renovate?

Renovate offers several advantages over the current Dependabot + Updatecli setup:

### Key Improvements

1. **Docker Security**: Automatically pins Docker base images with SHA256 digests
   - Current: `FROM debian` → Renovate: `FROM debian:latest@sha256:...`
   - Prevents supply chain attacks and ensures reproducible builds

2. **Intelligent Automerging**:
   - Patch updates: Auto-merge after 3 days stability
   - Digest updates: Auto-merge immediately
   - Security updates: Never auto-merge (requires manual review)

3. **Better Grouping Strategy**:
   - Groups dependencies by cloud provider (AWS, Azure, GCP)
   - Groups by framework (Elastic, K8s, Trivy)
   - Reduces PR noise from 50+ to ~10 PRs per week

4. **Vulnerability Dashboard**:
   - Centralized view of all pending updates
   - Priority scoring for security vulnerabilities
   - Immediate updates for critical/high severity issues

5. **Advanced Scheduling**:
   - Stability days: Wait 3 days before updating (avoid broken releases)
   - Minimum release age: Only update packages that have been released for 3+ days
   - Vulnerability alerts: Process immediately, any time

6. **Go Module Excellence**:
   - Automatic `go mod tidy` after updates
   - Updates import paths when packages move
   - Better handling of Go versioning (v2+, v7+)

## Configuration Details

### Key Features in `renovate.json`

```json
{
  "extends": [
    "config:recommended",        // Best practices
    ":dependencyDashboard",      // GitHub issue with all updates
    ":semanticCommits",          // Conventional commits
    ":separateMajorReleases"     // Major updates in separate PRs
  ]
}
```

### Package Rules Highlights

1. **Go Dependencies**:
   - Minor/patch: Grouped, weekly updates, 3-day stability period
   - Major: Separate PRs, lower priority
   - Grouped by provider: Azure, AWS, Google, Elastic, K8s, Trivy
   - Patch updates: Auto-merge after 3 days

2. **GitHub Actions**:
   - Weekly updates with digest pinning
   - Digest updates: Auto-merge immediately
   - Example: `actions/checkout@v4` → `actions/checkout@v4.1.1` (SHA pinned)

3. **Docker Images**:
   - Pins all images with SHA256 digests
   - Auto-merges digest updates
   - Security-focused approach

4. **Security Vulnerabilities**:
   - Highest priority (prPriority: 10)
   - Runs at any time (not limited to schedule)
   - Special labels: `security`, `priority`
   - Never auto-merged

### Custom Managers

The configuration includes a regex manager to track Go version updates in:
- `go.mod` file
- GitHub workflow files (`.github/workflows/*.yml`)

## Migration Steps

### Option 1: GitHub App (Recommended)

1. **Install Renovate GitHub App**:
   - Go to https://github.com/apps/renovate
   - Click "Install" or "Configure"
   - Select the `elastic/cloudbeat` repository

2. **Disable Dependabot** (after Renovate is working):
   - Remove or rename `.github/dependabot.yml`
   - Or keep it disabled in repo settings

3. **First Run**:
   - Renovate will create an "onboarding PR"
   - Review the PR to see what it will do
   - Merge the onboarding PR to activate

### Option 2: Self-Hosted

If you prefer to self-host Renovate:

1. **Create GitHub Workflow** (`.github/workflows/renovate.yml`):

```yaml
name: Renovate
on:
  schedule:
    - cron: '0 7 * * 1'  # Monday at 7am UTC
  workflow_dispatch:

jobs:
  renovate:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Self-hosted Renovate
        uses: renovatebot/github-action@v40.0.0
        with:
          configurationFile: renovate.json
          token: ${{ secrets.RENOVATE_TOKEN }}
        env:
          LOG_LEVEL: debug
```

2. **Create GitHub Token**:
   - Create a PAT with `repo` and `workflow` scopes
   - Add as secret: `RENOVATE_TOKEN`

## Comparison Matrix

| Feature | Dependabot | Renovate |
|---------|-----------|----------|
| Go modules | ✅ | ✅ |
| GitHub Actions | ✅ | ✅ (with digest pinning) |
| Docker images | ❌ | ✅ (with digest pinning) |
| Automerge | Limited | ✅ Advanced |
| Grouping | Basic | ✅ Advanced |
| Vulnerability dashboard | ❌ | ✅ |
| Stability days | ❌ | ✅ |
| Custom managers | ❌ | ✅ |
| Monorepo support | Limited | ✅ Excellent |
| Schedule flexibility | Basic | ✅ Advanced |

## What Changes After Migration

### Before (Dependabot)
```dockerfile
FROM debian
```

### After (Renovate)
```dockerfile
FROM debian:bookworm-slim@sha256:2b5263f1f2...
```

### PR Volume

**Current**:
- ~10-20 PRs per week (grouped)
- All require manual merge
- No security prioritization

**With Renovate**:
- ~5-10 PRs per week (better grouping)
- 30-50% auto-merged (patches, digests)
- Security updates clearly marked and prioritized

## Keeping Updatecli (Optional)

You can keep Updatecli for specific use cases:
- Updating from internal/private sources
- Complex update logic with conditionals
- Integration with Elastic's internal systems

Renovate handles public dependencies better, while Updatecli can focus on internal tooling.

## Testing the Configuration

Before going live:

1. **Fork Testing**:
   - Test on a fork first
   - Install Renovate app on your fork
   - Verify PR creation and automerge behavior

2. **Dry Run**:
   ```bash
   # Using Renovate CLI locally
   npm install -g renovate
   export GITHUB_TOKEN=your_token
   renovate --dry-run elastic/cloudbeat
   ```

3. **Monitor First Week**:
   - Check the Dependency Dashboard issue
   - Review automerge behavior
   - Adjust configuration if needed

## Rollback Plan

If issues arise:

1. **Disable Renovate App** (Settings → GitHub Apps)
2. **Re-enable Dependabot** (uncomment `.github/dependabot.yml`)
3. **Remove** `renovate.json`

## Support

- **Documentation**: https://docs.renovatebot.com/
- **Configuration**: https://docs.renovatebot.com/configuration-options/
- **Community**: https://github.com/renovatebot/renovate/discussions

## Next Steps

1. Review the `renovate.json` configuration
2. Choose installation method (GitHub App or self-hosted)
3. Test on a fork or feature branch first
4. Monitor the first week of operation
5. Adjust configuration based on team feedback
6. Consider deprecating Dependabot after successful migration
