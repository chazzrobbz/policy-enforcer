package policy_enforcer

type User struct {
	ID         string                 `json:"id"`
	Roles      []string               `json:"roles"`
	Attributes map[string]interface{} `json:"attributes"`
}

// NewUser .
func NewUser(id string, roles []string, attributes map[string]interface{}) User {
	return User{
		ID:         id,
		Roles:      roles,
		Attributes: attributes,
	}
}
