package policy_enforcer

import (
	`github.com/iancoleman/strcase`
)

// AnyOf */
func (p *Policy) AnyOf(any bool) (policy *Policy) {
	policy = p.getInstance()
	policy.Statement.AnyOf = any
	return
}

// Set */
func (p *Policy) Set(key string, value interface{}) (policy *Policy) {
	policy = p.getInstance()
	var mp map[string]interface{}
	mp, p.Error = ToMap(value)
	policy.Statement.Imports[key] = mp
	return
}

// Rule */
func (p *Policy) Rule(key string, conditions ...string) (policy *Policy) {
	policy = p.getInstance()
	var cn []string
	for _, con := range conditions {
		cn = append(cn, CleanCondition(con))
	}
	policy.Statement.Rules = append(policy.Statement.Rules, Rule{
		Key:        strcase.ToSnake(key),
		Conditions: cn,
	})
	return
}

// FailMessage */
func (p *Policy) FailMessage(key string, message string) (policy *Policy) {
	policy = p.getInstance()
	policy.Statement.Messages[strcase.ToSnake(key)] = message
	return
}
