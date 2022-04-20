package policy_enforcer

import (
	`context`
	`github.com/open-policy-agent/opa/storage`
)

// Statement statement
type Statement struct {
	Package string
	Imports map[string]interface{}
	Options []Option
	Storage storage.Store
	Rules   map[string]Rule
	Context context.Context
}
