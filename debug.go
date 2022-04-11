package policy_enforcer

import (
	`github.com/davecgh/go-spew/spew`
	`os`
)

// Pre exit running project.
// @param interface{}
// @param ...interface{}
func Pre(x interface{}, y ...interface{}) {
	spew.Dump(x)
	os.Exit(1)
}

