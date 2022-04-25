package policy_enforcer

import (
	"encoding/json"
	"math/rand"
	"strings"

	"github.com/gosimple/slug"
)

// CleanCondition weeds out the places where there might be errors when creating conditions
// @param string
// @return string
func CleanCondition(str string) string {
	return strings.ReplaceAll(str, "'", "\"")
}

// Key makes the keys you give while creating a rule compatible with the rego
// @param string
// @return string
func Key(b string) string {
	return strings.ReplaceAll(slug.Make(b), "-", "_")
}

// GenerateLowerCaseRandomString its generate lowercase random string
// @param int
// @return string
func GenerateLowerCaseRandomString(n int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

// ToMap turns objects into maps
// @param interface{}
// @return map[string]interface{}, error
func ToMap(u interface{}) (mp map[string]interface{}, err error) {
	var data []byte
	data, err = json.Marshal(u)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &mp)
	return
}

// ToMapArray turns objects into map array
// @param interface{}
// @return []map[string]interface{}, error
func ToMapArray(u interface{}) (mp []map[string]interface{}, err error) {
	var data []byte
	data, err = json.Marshal(u)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &mp)
	return
}
