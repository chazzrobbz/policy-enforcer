package policy_enforcer

import (
	`context`
)

// Policy */
type Policy struct {
	Error     error
	Statement *Statement
}

// New creates new policy instance
func New() *Policy {
	return &Policy{
		Error: nil,
		Statement: &Statement{
			Package: "app.permify",
			Imports: map[string]interface{}{},
			Options: []Option{},
			Rules:   map[string]Rule{},
			Storage: nil,
			Context: context.Background(),
		},
	}
}

// getInstance returns policy current instance
func (p *Policy) getInstance() *Policy {
	return p
}
