package policy_enforcer

import (
	`errors`
	`fmt`
	`strings`
)

type Rule struct {
	Key         string
	FailMessage error
	Conditions  []string
}

// NewRule creates a new rule
// @param string
// @param ...string
// @return Rule
func NewRule(conditions ...string) Rule {
	var cn []string
	for _, con := range conditions {
		cn = append(cn, CleanCondition(con))
	}
	return Rule{
		Key:        GenerateLowerCaseRandomString(20),
		Conditions: cn,
	}
}

// SetFailMessage set the message that it displays when the rule you have created return false
// @param string
// @return Rule
func (r Rule) SetFailMessage(message string) Rule {
	return Rule{
		Key:         r.Key,
		Conditions:  r.Conditions,
		FailMessage: errors.New(message),
	}
}

// SetKey set the key
// @param string
// @return Rule
func (r Rule) SetKey(key string) Rule {
	return Rule{
		Key:         Key(key),
		Conditions:  r.Conditions,
		FailMessage: r.FailMessage,
	}
}

// raw converts rule to string according to rego
func (r Rule) raw() string {
	return fmt.Sprintf(ruleTemplate, r.Key, strings.Join(r.Conditions, "\n"))
}

// Collection

// Rules provides methods for you to manage array data more easily.
type Rules []Rule

// Origin convert the collection to rule array.
// @return []models.Rule
func (c Rules) Origin() []Rule {
	return []Rule(c)
}

// Len returns the number of elements of the array.
// @return int64
func (c Rules) Len() (length int64) {
	return int64(len(c))
}

// Keys returns an array of the rule array's keys.
// @return []String
func (c Rules) Keys() (keys []string) {
	keys = []string{}
	for _, o := range c {
		keys = append(keys, o.Key)
	}
	return
}

// Raws returns an array of the rule array's raws.
// @return []String
func (c Rules) Raws() (raws []string) {
	raws = []string{}
	for _, o := range c {
		raws = append(raws, o.raw())
	}
	return
}
