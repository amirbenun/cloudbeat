package compliance.cis_k8s.rules.cis_5_1_6

import data.compliance.policy.kube_api.ensure_service_accounts as audit

finding = result {
	result := audit.finding(audit.service_account_automount)
}
