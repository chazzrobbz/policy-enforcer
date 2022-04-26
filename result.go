package policy_enforcer

// Result result */
type Result struct {
	index   int          `json:"-"`
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

// hasNext allows iterator */
func (r *Result) hasNext() bool {
	if r.index < len(r.Allows) {
		return true
	}
	return false
}

// getNext allows iterator */
func (r *Result) getNext() *Allow {
	if r.hasNext() {
		allow := r.Allows[r.index]
		r.index++
		return &allow
	}
	return nil
}
