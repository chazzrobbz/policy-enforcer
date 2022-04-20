package policy_enforcer

import (
	`fmt`
	`github.com/stretchr/testify/assert`
	`testing`
)

type User struct {
	Name   string   `json:"name"`
	Tenure int      `json:"tenure"`
	Roles  []string `json:"roles"`
}

func TestIsAuthorized_1(t *testing.T) {
	var user = User{
		Name:   "tolga",
		Tenure: 9,
		Roles:  []string{"admin"},
	}

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin").SetKey("is admin")
	var isSenior = NewRule("user.tenure > 8").SetFailMessage("user is not senior")
	var isManager = NewRule("'manager' in user.roles").SetFailMessage("user is not manager")

	policy := New()
	policy.Set("user", user)

	policy.Option(false, isAdmin).Option(false, isSenior, isManager)

	result, err := policy.IsAuthorized()

	fmt.Println(result)
	fmt.Println("=-=-=-=-=-=")
	Pre(result)

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, true)
}

func TestIsAuthorized_2(t *testing.T) {
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

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	var isSeniorManager = NewRule("user.tenure > 8", "'manager' in user.roles").SetFailMessage("user is not senior manager")

	policy.Option(false, isAdmin).Option(false, isSeniorManager)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, true)
}


func TestIsNotAuthorized_1(t *testing.T) {
	var user = struct {
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		Name:   "tolga",
		Tenure: 7,
		Roles:  []string{"manager"},
	}

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	var isSeniorManager = NewRule("user.tenure > 8", "'manager' in user.roles").SetFailMessage("user is not senior manager")

	policy := New()
	policy.Set("user", user)
	policy.Option(false, isAdmin).Option(false, isSeniorManager)

	fmt.Println(policy.ToRego())

	result, err := policy.IsAuthorized()

	fmt.Println(result)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, false)
}

func TestIsNotAuthorized_2(t *testing.T) {
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

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	var isSenior = NewRule("user.tenure > 8").SetFailMessage("user is not senior")
	var isManager = NewRule("'manager' in user.roles").SetFailMessage("user is not manager")

	policy.Option(false, isAdmin).Option(false, isSenior, isManager)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, false)
}

func TestIsAuthorizedWithResource_1(t *testing.T) {
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

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	var isSenior = NewRule("user.tenure > 8").SetFailMessage("user is not senior")
	var isManager = NewRule("'manager' in user.roles").SetFailMessage("user is not manager")
	var isResourceOwner = NewRule("resource.owner_id == user.id").SetFailMessage("user is not owner of the resource")

	policy := New()
	policy.Set("user", user)
	policy.Set("resource", resource)

	policy.Option(false, isAdmin).Option(false, isSenior, isManager).Option(false, isResourceOwner)

	result, err := policy.IsAuthorized()

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allow, true)
}
