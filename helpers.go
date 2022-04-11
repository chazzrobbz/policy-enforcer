package policy_enforcer

import (
	`encoding/json`
	`strings`
)

// CleanCondition */
func CleanCondition(str string) string {
	return strings.ReplaceAll(str, "'", "\"")
}

// ToMap */
func ToMap(u interface{}) (mp map[string]interface{}, err error)  {
	var data []byte
	data, err = json.Marshal(u)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &mp)
	return
}

