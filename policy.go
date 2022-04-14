package policy_enforcer

import (
	`context`
)

// Policy */
type Policy struct {
	Error     error
	Statement *Statement
}

// New */
func New() *Policy {
	return &Policy{
		Error: nil,
		Statement: &Statement{
			Package:  "app.permify",
			Imports:  map[string]interface{}{},
			Aliases:  map[string]string{},
			Messages: map[string]string{},
			Details:  map[string]map[string]interface{}{},
			Rules:    Rules{},
			AnyOf:    false,
			Context:  context.Background(),
		},
	}
}

// getInstance
func (p *Policy) getInstance() *Policy {
	return p
}
