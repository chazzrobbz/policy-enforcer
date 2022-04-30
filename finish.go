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

		var allowedResources []map[string]string

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
			Allow: r["allow"].(bool),
		}

		return Result{
			Allows:  []Allow{allow},
			Details: results,
		}, err

	}
}

// IsAuthorized It creates a decision about whether the inputs you give comply with the rules you write and a result about the reason
// @return Result, error
func (p *Policy) AuthorizedResources() (resources []Resource, err error) {
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

	if p.Statement.Strategy == MULTIPLE {
		var allowedResources []map[string]string

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

		return fill(p.Statement.Resources, allowedResources), nil
	} else {
		return []Resource{}, nil
	}
}

// ToRego Returns the rules you create programmatically as strings in rego language
// @return string
func (p *Policy) ToRego() string {
	var raw string

	for _, option := range p.Statement.Options {
		if Rules(option.Rules).Len() > 0 {
			raw += fmt.Sprintf(option.GetTemplate(p.Statement.Strategy), strings.Join(Rules(option.Rules).Heads(), "\n"))
		}
	}

	var defaults string
	if p.Statement.Strategy == SINGLE {
		defaults += fmt.Sprintf("default allow = false\n")
	}

	var imps string
	for _, imp := range p.Statement.Imports {
		imps += fmt.Sprintf("import input.%s as %s\n", imp, imp)
	}

	var rules []Rule
	for _, rule := range p.Statement.Rules {
		rules = append(rules, rule)
	}

	return fmt.Sprintf(policyTemplate, p.Statement.Package, imps, defaults, raw, strings.Join(Rules(rules).Evacuations(), ""))
}

// Compare */
func compare(allResources []Resource, allowedResourcesMap []map[string]string) (allows []Allow) {
	allowedMap := map[string]bool{}
	for _, r := range allowedResourcesMap {
		allowedMap[r["type"]+":"+r["id"]] = true
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

// Compare */
func fill(allResources []Resource, allowedResourcesMap []map[string]string) (resources []Resource) {
	allowedMap := map[string]bool{}
	for _, r := range allowedResourcesMap {
		allowedMap[r["type"]+":"+r["id"]] = true
	}
	for _, resource := range allResources {
		if allowedMap[resource.Type+":"+resource.ID] {
			resources = append(resources, resource)
		}
	}
	return
}
