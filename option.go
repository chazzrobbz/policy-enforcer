package policy_enforcer

// Option makes it easy for you to group rules and relate them to and,or
type Option struct {
	AnyOf bool
	Rules []Rule
}

// NewOption creates a new option
// @param bool
// @param ...Rule
// @return Option
func NewOption(anyOf bool, rules ...Rule) Option {
	return Option{
		AnyOf: anyOf,
		Rules: rules,
	}
}
