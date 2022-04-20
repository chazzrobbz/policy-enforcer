package policy_enforcer


//// ToPartials Returns the rules you create programmatically as strings in rego language
//// @return string
//func (p *Policy) ToPartials() (result []string, err error) {
//
//	r := rego.New(
//		rego.Query("data.app.permify.allow == true"),
//		rego.Module("permify.rego", p.ToRego()),
//	)
//
//	var pq rego.PreparedPartialQuery
//	pq, err = r.PrepareForPartial(p.Statement.Context)
//	if err != nil {
//		return
//	}
//
//	var pqs *rego.PartialQueries
//	pqs, err = pq.Partial(p.Statement.Context, rego.EvalInput(p.Statement.Imports), rego.EvalUnknowns([]string{"input.pets", "input.user"}))
//	if err != nil {
//		return
//	}
//
//	for i := range pqs.Queries {
//		result = append(result, fmt.Sprintf("%s", pqs.Queries[i]))
//	}
//
//	return
//}

// ToSQL Returns the rules you create programmatically as strings in rego language
// @return string
//func (p *Policy) ToSQL() string {
//	//p.ToPartials()
//
//
//
//}
