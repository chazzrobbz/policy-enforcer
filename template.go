package policy_enforcer

// policyTemplate main policy template
const policyTemplate = `
package %s

import future.keywords.every

# imports
%s

# options
%s

# rules
%s
`

// allowTemplate  main allow (option) template
const allowTemplate = `
allows[output] {
%s
output := {"allow": true}
}
`

const allowWithResourceTemplate = `
allows[output] {
resource := resources[_]
%s
output := {"id": resource.id, "type": resource.type, "allow": true}
}
`

// ruleTemplate  main rule template
const ruleTemplate = `
%s {
%s
}
`
