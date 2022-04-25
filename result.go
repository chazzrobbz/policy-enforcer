package policy_enforcer

// Result result */
type Result struct {
	Allows  []Allow      `json:"allows"`
	Details []RuleResult `json:"details"`
}

// Allow */
type Allow struct {
	Allow bool                   `json:"allow"`
	Meta  map[string]interface{} `json:"meta"`
}

// RuleResult rule result */
type RuleResult struct {
	Allow   bool   `json:"allow"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

