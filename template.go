package policy_enforcer

// policyTemplate main policy template
const policyTemplate = `
package %s

import future.keywords.every

# imports
%s

# defaults
%s

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

const allowWithResourceTemplate = `
allows[output] {
resource := resources[_]
%s
output := {"id": resource.id, "type": resource.type}
}
`

// ruleTemplate  main rule template
const ruleTemplate = `
%s {
%s
}
`

// validationTemplate
const validationTemplate = `
package validate

import future.keywords.every

default validate = false

validate {
	%s
}
`
