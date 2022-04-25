package policy_enforcer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/open-policy-agent/opa/rego"
)

// IsAuthorized It creates a decision about whether the inputs you give comply with the rules you write and a result about the reason
// @return Result, error
func (p *Policy) IsAuthorized() (result Result, err error) {
	var query rego.PreparedEvalQuery
	query, err = rego.New(
		rego.Query(fmt.Sprintf("r = data.%s", p.Statement.Package)),
		rego.Module("permify.rego", p.ToRego()),
	).PrepareForEval(p.Statement.Context)
	if err != nil {
		p.Error = err
		return
	}

	var resultSet rego.ResultSet
	resultSet, err = query.Eval(p.Statement.Context, rego.EvalInput(p.Statement.Inputs))
	if err != nil {
		p.Error = err
		return
	}

	r := resultSet[0].Bindings["r"].(map[string]interface{})

	var results []RuleResult

	for key, rule := range p.Statement.Rules {
		if !rule.ContainsResource {
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
	}

	if p.Statement.Strategy == MULTIPLE {
		var allowedResources []Resource

		var data []byte
		data, err = json.Marshal(r["allows"])
		if err != nil {
			p.Error = err
			return
		}

		err = json.Unmarshal(data, &allowedResources)
		if err != nil {
			p.Error = err
			return
		}

		return Result{
			Allows:  compare(p.Statement.Resources, allowedResources),
			Details: results,
		}, err
	} else {
		allow := Allow{
			Allow: true,
		}

		if len(r["allows"].([]interface{})) == 0 {
			allow.Allow = false
		}

		return Result{
			Allows:  []Allow{allow},
			Details: results,
		}, err
	}
}

// ToRego Returns the rules you create programmatically as strings in rego language
// @return string
func (p *Policy) ToRego() string {
	var raw string

	for _, option := range p.Statement.Options {
		if Rules(option.Rules).Len() > 0 {
			raw += fmt.Sprintf(option.GetTemplate(p.Statement.Strategy), strings.Join(Rules(option.Rules).Titles(), "\n"))
		}
	}

	var imps string
	for _, imp := range p.Statement.Imports {
		imps += fmt.Sprintf("import input.%s as %s\n", imp, imp)
	}

	var rules []Rule
	for _, rule := range p.Statement.Rules {
		rules = append(rules, rule)
	}

	return fmt.Sprintf(policyTemplate, p.Statement.Package, imps, raw, strings.Join(Rules(rules).Evacuations(), ""))
}

// Compare */
func compare(allResources []Resource, allowedResources []Resource) (allows []Allow) {
	allowedMap := map[string]bool{}
	for _, allowedResource := range allowedResources {
		allowedMap[allowedResource.Type+":"+allowedResource.ID] = true
	}
	for _, resource := range allResources {
		if allowedMap[resource.Type+":"+resource.ID] {
			allows = append(allows, Allow{
				Allow: true,
				Meta: map[string]interface{}{
					"id":   resource.ID,
					"type": resource.Type,
				},
			})
		} else {
			allows = append(allows, Allow{
				Allow: false,
				Meta: map[string]interface{}{
					"id":   resource.ID,
					"type": resource.Type,
				},
			})
		}
	}
	return
}
