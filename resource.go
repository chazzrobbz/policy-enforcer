package policy_enforcer

type Resource struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

// NewResource .
func NewResource(id string, typ string, attributes map[string]interface{}) Resource {
	return Resource{
		ID:         id,
		Type:       typ,
		Attributes: attributes,
	}
}
