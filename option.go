package policy_enforcer

// Option makes it easy for you to group rules and relate them to and,or
type Option struct {
	Rules []Rule
}

// NewOption creates a new option
// @param bool
// @param ...Rule
// @return Option
func NewOption(rules ...Rule) Option {
	return Option{
		Rules: rules,
	}
}

// GetTemplate
// @param string
// @return string
func (r Option) GetTemplate(s Strategy) string {
	switch s {
	case MULTIPLE:
		return allowWithResourceTemplate
	case SINGLE:
		return allowTemplate
	default:
		return allowTemplate
	}
}
