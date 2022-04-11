package policy_enforcer


import (
	`context`
)

// Statement */
type Statement struct {
	Package  string
	Imports  map[string]interface{}
	Messages map[string]string
	Rules    []Rule
	AnyOf    bool
	Context  context.Context
}

