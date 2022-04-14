package policy_enforcer

import (
	`context`
)

// Statement */
type Statement struct {
	Package  string
	Imports  map[string]interface{}
	Aliases  map[string]string
	Messages map[string]string
	Details  map[string]map[string]interface{}
	Rules    []Rule
	AnyOf    bool
	Context  context.Context
}
