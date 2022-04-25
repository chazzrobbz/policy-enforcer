package policy_enforcer

import (
	`github.com/stretchr/testify/assert`
	`testing`
)

func Test1(t *testing.T) {
	var user = User{
		ID: "tolga",
		Attributes: map[string]interface{}{
			"tenure": 8,
		},
		Roles: []string{"admin"},
	}

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin").SetKey("is admin")
	var isSenior = NewRule("user.attributes.tenure > 8").SetFailMessage("user is not senior")
	var isManager = NewRule("'manager' in user.roles").SetFailMessage("user is not manager")

	policy := New()
	policy.SetUser(user)

	policy.Option(isAdmin).Option(isSenior, isManager)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, true)
}

func Test2(t *testing.T) {

	policy := New()

	policy.SetUser(User{
		ID:    "1",
		Roles: []string{"manager"},
		Attributes: map[string]interface{}{
			"tenure": 9,
		},
	})

	policy.SetResources(
		Resource{
			ID:   "1",
			Type: "posts",
			Attributes: map[string]interface{}{
				"owner_id": "1",
			},
		},
		Resource{
			ID:   "2",
			Type: "posts",
			Attributes: map[string]interface{}{
				"owner_id": "2",
			},
		},
	)

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	var isResourceOwner = NewRule("resource.attributes.owner_id == '1'")

	policy.Option(isAdmin).Option(isResourceOwner)

	var r, err = policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, r.Allows[0].Allow, true)
	assert.Equal(t, r.Allows[1].Allow, false)
}

func Test3(t *testing.T) {
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
	policy.Option(isAdmin).Option(isSeniorManager)

	result, err := policy.IsAuthorized()

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, false)
}

func Test4(t *testing.T) {
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

	policy.Option(isAdmin).Option(isSenior, isManager)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, false)
}

func Test5(t *testing.T) {
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

	var post = struct {
		ID      string `json:"id"`
		OwnerID int    `json:"owner_id"`
	}{
		ID:      "1",
		OwnerID: 1,
	}

	var isAdmin = NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	var isSenior = NewRule("user.tenure > 8").SetFailMessage("user is not senior")
	var isManager = NewRule("'manager' in user.roles").SetFailMessage("user is not manager")
	var isResourceOwner = NewRule("post.owner_id == user.id").SetFailMessage("user is not owner of the post")

	policy := New()
	policy.Set("user", user)
	policy.Set("post", post)

	policy.Option(isAdmin).Option(isSenior, isManager).Option(isResourceOwner)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, true)
}
