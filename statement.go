package policy_enforcer

import (
	`context`
)

// Statement statement
type Statement struct {
	Package   string
	Imports   []string
	Inputs    map[string]interface{}
	Options   []Option
	Rules     map[string]Rule
	User      User
	Resources []Resource
	Strategy  Strategy
	Context   context.Context
}
