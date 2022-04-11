package policy_enforcer

import (
	"github.com/stretchr/testify/assert"
	`testing`
)

func TestIsAuthorized_AnyOf(t *testing.T) {
	var user = struct {
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		Name:   "tolga",
		Tenure: 9,
		Roles:  []string{"admin"},
	}

	policy := New()
	policy.Set("user", user)
	policy.AnyOf(true)

	policy.Rule("is admin", "'admin' in user.roles")
	policy.FailMessage("is admin", "user is not an admin")

	policy.Rule("is senior manager", "user.tenure > 8", "'manager' in user.roles")
	policy.FailMessage("is senior manager", "user is not senior manager")

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, true)
}

func TestIsNotAuthorized_AnyOf(t *testing.T) {
	var user = struct {
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		Name:   "tolga",
		Tenure: 7,
		Roles:  []string{"manager"},
	}

	policy := New()
	policy.Set("user", user)
	policy.AnyOf(true)

	policy.Rule("is admin", "'admin' in user.roles")
	policy.FailMessage("is admin", "user is not an admin")

	policy.Rule("is senior manager", "user.tenure > 8", "'manager' in user.roles")
	policy.FailMessage("is senior manager", "user is not senior manager")

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, false)
}

func TestIsAuthorized(t *testing.T) {
	var user = struct {
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		Name:   "tolga",
		Tenure: 9,
		Roles:  []string{"admin", "manager"},
	}

	policy := New()
	policy.Set("user", user)
	policy.AnyOf(false)

	policy.Rule("is admin", "'admin' in user.roles")
	policy.FailMessage("is admin", "user is not an admin")

	policy.Rule("is senior manager", "user.tenure > 8", "'manager' in user.roles")
	policy.FailMessage("is senior manager", "user is not senior manager")

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, true)
}

func TestIsNotAuthorized(t *testing.T) {
	var user = struct {
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		Name:   "tolga",
		Tenure: 7,
		Roles:  []string{"admin", "manager"},
	}

	policy := New()
	policy.Set("user", user)
	policy.AnyOf(false)

	policy.Rule("is admin", "'admin' in user.roles")
	policy.FailMessage("is admin", "user is not an admin")

	policy.Rule("is senior manager", "user.tenure > 8", "'manager' in user.roles")
	policy.FailMessage("is senior manager", "user is not senior manager")

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, false)
}

func TestIsAuthorizedWithResource_AnyOf(t *testing.T) {
	var user = struct {
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		ID:     1,
		Name:   "tolga",
		Tenure: 9,
		Roles:  []string{"admin"},
	}

	var resource = struct {
		Type    string `json:"type"`
		OwnerID int    `json:"owner_id"`
	}{
		Type:    "posts",
		OwnerID: 1,
	}

	policy := New()
	policy.Set("user", user)
	policy.Set("resource", resource)
	policy.AnyOf(true)

	policy.Rule("is admin", "'admin' in user.roles")
	policy.FailMessage("is admin", "user is not an admin")

	policy.Rule("is senior manager", "user.tenure > 8", "'manager' in user.roles")
	policy.FailMessage("is senior manager", "user is not senior manager")

	policy.Rule("is resource owner", "resource.owner_id == user.id")
	policy.FailMessage("is resource owner", "user is not owner of the resource")

	result, err := policy.IsAuthorized()

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, true)
}

