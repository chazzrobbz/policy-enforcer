package policy_enforcer

// Result */
type Result struct {
	Allow   bool         `json:"allow"`
	Details []RuleResult `json:"details"`
}

// RuleResult */
type RuleResult struct {
	Allow   bool                   `json:"allow"`
	Key     string                 `json:"key"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}
