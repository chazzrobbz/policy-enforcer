package policy_enforcer

// policyTemplate main policy template
const policyTemplate = `
package %s

import future.keywords.every

# imports
%s
default allow = false

# options
%s
# rules
%s
`

// allowTemplate  main allow (option) template
const allowTemplate = `
allow {
%s
}
`

// ruleTemplate  main rule template
const ruleTemplate = `
%s {
%s
}
`
