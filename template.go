package policy_enforcer

// policyTemplate */
const policyTemplate = `
package %s

import future.keywords.every

# imports
%s
default allow = false
%s
# rules
%s
`

// allowTemplate */
const allowTemplate = `
allow {
%s
}
`

// ruleTemplate */
const ruleTemplate = `
%s {
%s
}
`