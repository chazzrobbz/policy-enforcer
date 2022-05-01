package policy_enforcer

// Set loads the object into the statement.
// @param string
// @param interface{}
// @return *Policy
func (p *Policy) Set(key string, value interface{}) (policy *Policy) {
	policy = p.getInstance()
	var mp map[string]interface{}
	mp, p.Error = ToMap(value)
	policy.Statement.Imports = append(policy.Statement.Imports, key)
	policy.Statement.Inputs[key] = mp
	return
}

// SetUser loads the user into the statement.
// @param string
// @param interface{}
// @return *Policy
func (p *Policy) SetUser(user User) (policy *Policy) {
	policy = p.getInstance()
	policy.Statement.User = user
	var mp map[string]interface{}
	mp, p.Error = ToMap(user)
	policy.Statement.Imports = append(policy.Statement.Imports, "user")
	policy.Statement.Inputs["user"] = mp
	return
}

// SetResources loads the resources into the statement.
// @param string
// @param interface{}
// @return *Policy
func (p *Policy) SetResources(resources ...Resource) (policy *Policy) {
	if len(resources) > 0 {
		policy = p.getInstance()
		policy.Statement.Resources = resources
		var mp []map[string]interface{}
		mp, p.Error = ToMapArray(resources)
		policy.Statement.Imports = append(policy.Statement.Imports, "resources")
		policy.Statement.Inputs["resources"] = mp
		p.setStrategy(MULTIPLE)
	}
	return
}

// Option makes it easy for you to group rules and relate them to and
// @param bool
// @param ...Rule
// @return *Policy
func (p *Policy) Option(rules ...Rule) (policy *Policy) {
	if len(rules) > 0 {
		policy = p.getInstance()
		policy.Statement.Options = append(policy.Statement.Options, Option{
			Rules: rules,
		})
		for _, rule := range rules {
			policy.Statement.Rules[rule.Key] = rule
		}
	}
	return
}

// SetPackage .
// @param string
// @return *Policy
func (p *Policy) SetPackage(pg string) (policy *Policy) {
	policy = p.getInstance()
	policy.Statement.Package = pg
	return
}

// setStrategy .
// @param Strategy
// @return *Policy
func (p *Policy) setStrategy(s Strategy) (policy *Policy) {
	policy = p.getInstance()
	policy.Statement.Strategy = s
	return
}
