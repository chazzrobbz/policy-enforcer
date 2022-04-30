package policy_enforcer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	user := User{
		ID: "1",
		Attributes: map[string]interface{}{
			"tenure": 3,
		},
		Roles: []string{"admin"},
	}

	blog := map[string]interface{}{
		"id":     1,
		"status": "PUBLIC",
	}

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin").SetKey("is admin")
	isSenior := NewRule("user.attributes.tenure > 8").SetFailMessage("user is not senior")
	isManager := NewRule("'manager' in user.roles").SetFailMessage("user is not manager")
	isPublic := NewRule("blog.status == 'PUBLIC'").SetFailMessage("blog is not public")

	policy := New()
	policy.SetUser(user)
	policy.Set("blog", blog)

	policy.Option(isAdmin).Option(isSenior, isManager).Option(isPublic)

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

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	isResourceOwner := NewRule("resource.attributes.owner_id == '1'")

	policy.Option(isAdmin).Option(isResourceOwner)

	result, err := policy.IsAuthorized()

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, true)
	assert.Equal(t, result.Allows[1].Allow, false)
}

func Test3(t *testing.T) {

	policy := New()

	policy.SetUser(User{
		ID:    "1",
		Roles: []string{"admin"},
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

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	isResourceOwner := NewRule("resource.attributes.owner_id == '1'")

	policy.Option(isAdmin).Option(isResourceOwner)

	result, err := policy.IsAuthorized()

	assert.Equal(t, err, nil)

	for result.hasNext() {
		allow := result.getNext()
		assert.Equal(t, allow.Allow, true)
	}
}

func Test4(t *testing.T) {
	user := struct {
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		Name:   "tolga",
		Tenure: 7,
		Roles:  []string{"manager"},
	}

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	isSeniorManager := NewRule("user.tenure > 8", "'manager' in user.roles").SetFailMessage("user is not senior manager")

	policy := New()
	policy.Set("user", user)
	policy.Option(isAdmin).Option(isSeniorManager)

	result, err := policy.IsAuthorized()

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, false)
}

func Test5(t *testing.T) {
	user := struct {
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

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	isSenior := NewRule("user.tenure > 8").SetFailMessage("user is not senior")
	isManager := NewRule("'manager' in user.roles").SetFailMessage("user is not manager")

	policy.Option(isAdmin).Option(isSenior, isManager)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, false)
}

func Test6(t *testing.T) {
	user := struct {
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Tenure int      `json:"tenure"`
		Roles  []string `json:"roles"`
	}{
		ID:     1,
		Name:   "tolga",
		Tenure: 9,
		Roles:  []string{},
	}

	post := struct {
		ID      string `json:"id"`
		OwnerID int    `json:"owner_id"`
	}{
		ID:      "1",
		OwnerID: 1,
	}

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	isResourceOwner := NewRule("post.owner_id == user.id").SetFailMessage("user is not owner of the post")

	policy := New()
	policy.Set("user", user)
	policy.Set("post", post)

	policy.Option(isAdmin).Option(isResourceOwner)

	result, err := policy.IsAuthorized()
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Allows[0].Allow, true)
}

func Test7(t *testing.T) {

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

	isAdmin := NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
	isResourceOwner := NewRule("resource.attributes.owner_id == '1'")

	policy.Option(isAdmin).Option(isResourceOwner)

	resources, err := policy.AuthorizedResources()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(resources), 1)
}

func Test8(t *testing.T) {

	var err error
	isAdmin := NewRule("'admin' in user.roles")
	isResourceOwner := NewRule("resource.attributes.owner_id == '1'")

	err = isAdmin.Validate()
	assert.Equal(t, err, nil)

	err = isResourceOwner.Validate()
	assert.Equal(t, err, nil)
}

func Test9(t *testing.T) {
	err := NewRule("'admin i user.roles").Validate()
	assert.NotEmpty(t, err)
}
