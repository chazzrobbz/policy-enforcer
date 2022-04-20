package policy_enforcer

// Set loads the object into the statement.
// @param string
// @param interface{}
// @return *Policy
func (p *Policy) Set(key string, value interface{}) (policy *Policy) {
	policy = p.getInstance()
	var mp map[string]interface{}
	mp, p.Error = ToMap(value)
	policy.Statement.Imports[key] = mp
	return
}

// Option makes it easy for you to group rules and relate them to and,or
// @param bool
// @param ...Rule
// @return *Policy
func (p *Policy) Option(anyOf bool, rules ...Rule) (policy *Policy) {
	policy = p.getInstance()
	policy.Statement.Options = append(policy.Statement.Options, Option{
		AnyOf: anyOf,
		Rules: rules,
	})
	for _, rule := range rules {
		policy.Statement.Rules[rule.Key] = rule
	}
	return
}
