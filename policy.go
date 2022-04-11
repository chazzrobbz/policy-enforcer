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
			Messages: map[string]string{},
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

