package policy_enforcer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/open-policy-agent/opa/rego"
)

type Rule struct {
	Key              string
	ContainsResource bool
	FailMessage      error
	Conditions       []string
}

// NewRule creates a new rule
// @param string
// @param ...string
// @return Rule
func NewRule(conditions ...string) Rule {
	var cn []string
	containsResource := false
	for _, con := range conditions {
		condition := CleanCondition(con)
		containsResource = strings.Contains(condition, "resource.")
		cn = append(cn, condition)
	}
	return Rule{
		Key:              GenerateLowerCaseRandomString(20),
		ContainsResource: containsResource,
		Conditions:       cn,
	}
}

// SetFailMessage set the message that it displays when the rule you have created return false
// @param string
// @return Rule
func (r Rule) SetFailMessage(message string) Rule {
	return Rule{
		Key:              r.Key,
		Conditions:       r.Conditions,
		FailMessage:      errors.New(message),
		ContainsResource: r.ContainsResource,
	}
}

// SetKey set the key
// @param string
// @return Rule
func (r Rule) SetKey(key string) Rule {
	return Rule{
		Key:              Key(key),
		Conditions:       r.Conditions,
		FailMessage:      r.FailMessage,
		ContainsResource: r.ContainsResource,
	}
}

// GetHead
// @param string
// @return string
func (r Rule) GetHead() string {
	if r.ContainsResource {
		return fmt.Sprintf("%s(resource)", r.Key)
	}
	return r.Key
}

// GetHead
// @param string
// @return string
func (r Rule) GetBody() string {
	return strings.Join(r.Conditions, "\n")
}

// GetTemplate
// @param string
// @return string
func (r Rule) GetTemplate() string {
	return ruleTemplate
}

// Validate
// @return error
func (r Rule) Validate() error {
	_, err := rego.New(
		rego.Query("r = data.validate"),
		rego.Module("permify.rego", fmt.Sprintf(validationTemplate, r.GetBody())),
	).PrepareForEval(nil)
	if err != nil {
		if strings.Contains(err.Error(), "rego_unsafe_var_error") {
			return nil
		}
		return err
	}
	return nil
}

// Evict
// @param string
// @return string
func (r Rule) Evict() string {
	return fmt.Sprintf(r.GetTemplate(), r.GetHead(), r.GetBody())
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

// Titles returns an array of the rule array's keys.
// @return []String
func (c Rules) Heads() (keys []string) {
	keys = []string{}
	for _, o := range c {
		keys = append(keys, o.GetHead())
	}
	return
}

// Evacuations returns an array of the rule array's raws.
// @return []String
func (c Rules) Evacuations() (evacuations []string) {
	evacuations = []string{}
	for _, o := range c {
		evacuations = append(evacuations, o.Evict())
	}
	return
}
