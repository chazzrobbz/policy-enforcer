package policy_enforcer

import (
	`fmt`
	`github.com/open-policy-agent/opa/rego`
	`strings`
)

// IsAuthorized */
func (p *Policy) IsAuthorized() (result Result, err error) {
	var query rego.PreparedEvalQuery
	query, err = rego.New(
		rego.Query(fmt.Sprintf("r = data.app.permify")),
		rego.Module("permify.rego", p.ToRego()),
	).PrepareForEval(p.Statement.Context)
	if err != nil {
		return
	}

	var resultSet rego.ResultSet
	resultSet, err = query.Eval(p.Statement.Context, rego.EvalInput(p.Statement.Imports))
	if err != nil {
		return
	}

	r := resultSet[0].Bindings["r"].(map[string]interface{})

	var results []RuleResult
	for _, rule := range Rules(p.Statement.Rules) {
		_, ok := r[rule.Key].(bool)
		if ok {
			results = append(results, RuleResult{
				Allow:   true,
				Key:     rule.Key,
			})
		} else {
			results = append(results, RuleResult{
				Allow:   false,
				Key:     rule.Key,
				Message: p.Statement.Messages[rule.Key],
			})
		}
	}

	return Result{
		Allow:   r["allow"].(bool),
		Details: results,
	}, err
}

// ToRego */
func (p *Policy) ToRego() string {
	var rv string
	if p.Statement.AnyOf {
		for _, name := range Rules(p.Statement.Rules).Keys() {
			rv += fmt.Sprintf(allowTemplate, name)
		}
	} else {
		rv = fmt.Sprintf(allowTemplate, strings.Join(Rules(p.Statement.Rules).Keys(), "\n"))
	}
	var imps string
	for key, _ := range p.Statement.Imports {
		imps += fmt.Sprintf("import input.%s as %s\n", key, key)
	}
	return fmt.Sprintf(policyTemplate, p.Statement.Package, imps, rv, strings.Join(Rules(p.Statement.Rules).Raws(), ""))
}
