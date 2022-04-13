package policy_enforcer

import (
	`encoding/json`
	`github.com/gosimple/slug`
	`strings`
)

// CleanCondition */
func CleanCondition(str string) string {
	return strings.ReplaceAll(str, "'", "\"")
}

// Key */
func Key(b string) string {
	return strings.ReplaceAll(slug.Make(b), "-", "_")
}

// ToMap */
func ToMap(u interface{}) (mp map[string]interface{}, err error) {
	var data []byte
	data, err = json.Marshal(u)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &mp)
	return
}
