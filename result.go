package policy_enforcer

// Result result */
type Result struct {
	Allow   bool         `json:"allow"`
	Details []RuleResult `json:"details"`
}

// RuleResult rule result */
type RuleResult struct {
	Allow   bool                   `json:"allow"`
	Key     string                 `json:"key"`
	Message string                 `json:"message"`
}
