package policy_enforcer

// Result result */
type Result struct {
	Allows  []Allow      `json:"allows"`
	Details []RuleResult `json:"details"`
}

type Allow struct {
	Allow bool                   `json:"allow"`
	Meta  map[string]interface{} `json:"meta"`
}

// RuleResult rule result */
type RuleResult struct {
	Allow   bool   `json:"allow"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

func Compare(allResources []Resource, allowedResources []Resource) (allows []Allow) {
	var allowedMap = map[string]bool{}
	for _, allowedResource := range allowedResources {
		allowedMap[allowedResource.Type+":"+allowedResource.ID] = true
	}
	for _, resource := range allResources {
		if allowedMap[resource.Type+":"+resource.ID] {
			allows = append(allows, Allow{
				Allow: true,
				Meta: map[string]interface{}{
					"id":   resource.ID,
					"type": resource.Type,
				},
			})
		} else {
			allows = append(allows, Allow{
				Allow: false,
				Meta: map[string]interface{}{
					"id":   resource.ID,
					"type": resource.Type,
				},
			})
		}
	}
	return
}
