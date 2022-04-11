package policy_enforcer

import (
	`fmt`
	`strings`
)

// Rule */
type Rule struct {
	Key       string
	Conditions []string
}

// raw */
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
