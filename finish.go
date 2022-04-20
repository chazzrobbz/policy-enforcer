package policy_enforcer

import (
	`fmt`
	`github.com/open-policy-agent/opa/rego`
	`strings`
)

// IsAuthorized It creates a decision about whether the inputs you give comply with the rules you write and a result about the reason
// @return Result, error
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
	for key, rule := range p.Statement.Rules {
		_, ok := r[rule.Key].(bool)
		if ok {
			results = append(results, RuleResult{
				Allow:   true,
				Key:     key,
				Message: "",
			})
		} else {
			results = append(results, RuleResult{
				Allow:   false,
				Key:     key,
				Message: rule.FailMessage.Error(),
			})
		}
	}

	return Result{
		Allow:   r["allow"].(bool),
		Details: results,
	}, err
}

// ToRego Returns the rules you create programmatically as strings in rego language
// @return string
func (p *Policy) ToRego() string {
	var raw string

	for _, option := range p.Statement.Options {
		if option.AnyOf {
			for _, name := range Rules(option.Rules).Keys() {
				raw += fmt.Sprintf(allowTemplate, name)
			}
		} else {
			if Rules(option.Rules).Len() > 0 {
				raw += fmt.Sprintf(allowTemplate, strings.Join(Rules(option.Rules).Keys(), "\n"))
			}
		}
	}

	var imps string
	for key, _ := range p.Statement.Imports {
		imps += fmt.Sprintf("import input.%s as %s\n", key, key)
	}

	var rules []Rule
	for _, rule := range p.Statement.Rules {
		rules = append(rules, rule)
	}

	return fmt.Sprintf(policyTemplate, p.Statement.Package, imps, raw, strings.Join(Rules(rules).Raws(), ""))
}