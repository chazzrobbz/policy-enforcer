
<h1 align="center">
  <img alt="policy enforcer logo" src="https://user-images.githubusercontent.com/39353278/165170342-d1d61da0-f464-48c5-a3df-6a9bf1d14aa2.png" width="250px"/><br/>
  Policy Enforcer
</h1>

<p align="center">Represent your rego rules programmatically.</p>

<p align="center"><a href="https://pkg.go.dev/github.com/Permify/policy-enforcer?tab=doc" 
target="_blank"></a><img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;&nbsp;<img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" />&nbsp;&nbsp;<img src="https://img.shields.io/github/license/Permify/policy-enforcer?style=for-the-badge" alt="license" />&nbsp;&nbsp;<img src="https://img.shields.io/github/last-commit/Permify/policy-enforcer?style=for-the-badge" alt="tweet" />&nbsp;&nbsp;<img src="https://img.shields.io/twitter/url?style=for-the-badge&url=https%3A%2F%2Ftwitter.com%2Fgetpermify" alt="tweet" /></p>


Policy enforcer is a open source tool that allows you to easily create complex rego.

> Rego is the policy language for defining rules that are evaluated by the OPA (Open Policy Agent) engine.

## Features

- Generate your complex authorization easily with code.
- Export the rego you created with the code.
- Make decisions about multiple resources from one policy.
- Get the details of the decisions made.
- Add custom messages and handle decision messages.

## 👇 Setup

Install

```shell
go get github.com/Permify/policy-enforcer
```

Run Tests

```shell
go test
```

Import enforcer.

```go
import enforcer `github.com/Permify/policy-enforcer`
```

## 🚀 Usage

```go
var user = User{
    ID:   1,
    Roles:  []string{"admin"},
    Attributes: map[string]interface{}{
            "tenure": 8,
    },
}

blog := map[string]interface{}{
    "id":     1,
    "status": "PUBLIC",
}

var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin").SetKey("is admin")
var isSenior = enforcer.NewRule("user.tenure > 8").SetFailMessage("user is not senior")
var isManager = enforcer.NewRule("'manager' in user.roles").SetFailMessage("user is not manager")
var isPublic = enforcer.NewRule("blog.status == 'PUBLIC'").SetFailMessage("blog is not public")

policy := enforcer.New()
    
// set user object
policy.Set("user", user)
policy.Set("blog", blog)

// its means the user must be either an admin or a senior manager or blog is public
policy.Option(isAdmin).Option(isSenior, isManager).Option(isPublic)

result, err := policy.IsAuthorized()
```

### Output
```go
{
    Allows: {
        {
            Allow: true // final result
        }, 	
    },
    Details: {
        {
            Allow: true, // is admin result
            Key: "is_admin", 
            Message: ""
        },
        {
            Allow: true,  // is senior result
            Key: "tcuaxhxkqfdafplsjfbc", // if the key is not set it is created automatically
            Message: ""
        },
        {
            Allow: true, // is public result
            Key: (string) (len=20) "lgtemapezqleqyhyzryw",
            Message: (string) ""
        },
        {
            Allow: false, // is manager result
            Key: "xoeffrswxpldnjobcsnv", // if the key is not set it is created automatically
            Message: "user is not manager"
        }
    }
}
```

## 🔗 Multiple Resource Response

You can check authorization of multiple resources at once.

example scenario
- If the user is admin, can edit all listed resources. (Admin)
- If the user owns the resource, they can only edit their own resources. (Resource Owner)

### Admin

```go
policy := enforcer.New()

policy.SetUser(enforcer.User{
    ID:    "1",
    Roles: []string{"admin"}, // admin
    Attributes: map[string]interface{}{
        "tenure": 9,
    },
})

policy.SetResources(
    enforcer.Resource{
        ID:   "1",
        Type: "posts",
        Attributes: map[string]interface{}{
            "owner_id": "1",
        },
    },
    enforcer.Resource{
        ID:   "2",
        Type: "posts",
        Attributes: map[string]interface{}{
            "owner_id": "2",
        },
    },
)

var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
var isResourceOwner = enforcer.NewRule("resource.attributes.owner_id == '1'")

// its means the user must be either an admin or a resource owner
policy.Option(isAdmin).Option(isResourceOwner)

var r, err = policy.IsAuthorized()
```

```go
{
    Allows: {
        {
            Allow: true, // its true because user is admin
            Meta: {
                 "id": "1"
				 "type": "posts",
            }
        },
        {
            Allow: true, // its true because user is admin
            Meta: {
                "id": "2",
                "type": "posts"
            }
        }
    },
    Details: {
        {
            Allow: true,
            Key: "lgtemapezqleqyhyzryw",
            Message: ""
        }
    }
}
```

### Resource Owner

```go
policy := New()

policy.SetUser(enforcer.User{
    ID:    "1",
    Roles: []string{"manager"},
    Attributes: map[string]interface{}{
        "tenure": 9,
    },
})

policy.SetResources(
    enforcer.Resource{
        ID:   "1",
        Type: "posts",
        Attributes: map[string]interface{}{
            "owner_id": "1",
        },
    },
    enforcer.Resource{
        ID:   "2",
        Type: "posts",
        Attributes: map[string]interface{}{
            "owner_id": "2",
        },
    },
)

var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
var isResourceOwner = enforcer.NewRule("resource.attributes.owner_id == '1'")

// its means the user must be either an admin or a resource owner
policy.Option(isAdmin).Option(isResourceOwner)

var r, err = policy.IsAuthorized()
```

### Output

```go
{
    Allows: {
        {
            Allow: true, // its true because user is owner of the this resource
            Meta: {
		        "id": "1"
                "type": "posts",
            }
        },
        {
            Allow: false,
            Meta: {
                "id": "2",
                "type": "posts"
            }
        }
    },
    Details: {
        {
            Allow: false,
            Key: "lgtemapezqleqyhyzryw",
            Message: "user is not an admin"
        }
    }
}
```


## 🚨 Create New Rule

the user should a manager role among their roles
```go
var isManager = NewRule("'manager' in user.roles")
```

The user's tenure must be at least 8 years ***and*** the user should a manager role among their roles.
```go
var isSeniorManager = NewRule("user.attributes.tenure > 8", "'manager' in user.roles")
```

The user is the owner of the resource or resources.
```go
var isResourceOwner = NewRule("resource.attributes.owner.id == '1'")
```

### ⁉️ Set Fail Message

After the set fail message function decides on the policy, if the rule is false, this message will be printed on the error

```go
var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
```

### Output

```go
Details: {
    {
        Allow: false, // result
        Key: "xoeffrswxpldnjobcsnv",
        Message: "user is not an admin" // when it fails the message will appear here
    },
}
```

### 🔑 Set Key

You can use it when you do not want the key to be generated automatically. This will allow you to perceive your details better.

```go
var isAdmin = enforcer.NewRule("'admin' in user.roles").SetKey("is admin").SetFailMessage("user is not an admin")
```

### Output

```go
Details: {
    {
        Allow: true, // result
        Key: "is_admin",
        Message: "" // when true, the message does not appear
    },
}
```

## Options

Options allow you to establish an ***or*** relationship between rules and create a more organized and legible authorization structure.

```go
// its means the user must be either an admin or a senior and manager
enforcer.Option(isAdmin).Option(isSenior, isManager)
```


## To Rego Function

You can export the rego policies you created with the code with this function.

### Example 1

```go
var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin").SetKey("is admin")
var isSenior = enforcer.NewRule("user.tenure > 8").SetFailMessage("user is not senior")
var isManager = enforcer.NewRule("'manager' in user.roles").SetFailMessage("user is not manager")

policy := enforcer.New()

// its means the user must be either an admin or a senior manager
policy.Option(isAdmin).Option(isSenior, isManager)

fmt.Println(policy.ToRego())
```

### Output
```
package app.permify

import future.keywords.every

# imports
import input.user as user

# options
allows[output] {
  is_admin
  output := {"allow": true}
}

allows[output] {
  tcuaxhxkqfdafplsjfbc
  xoeffrswxpldnjobcsnv
  output := {"allow": true}
}

# rules

tcuaxhxkqfdafplsjfbc {
  user.attributes.tenure > 8
}

xoeffrswxpldnjobcsnv {
  "manager" in user.roles
}

is_admin {
  "admin" in user.roles
}
```

### Example 2

```go
policy := enforcer.New()

var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin")
var isResourceOwner = enforcer.NewRule("resource.attributes.owner_id == '1'")

// its means the user must be either an admin or a resource owner
policy.Option(isAdmin).Option(isResourceOwner)

fmt.Println(policy.ToRego())
```

### Output

```
package app.permify

import future.keywords.every

# imports
import input.user as user
import input.resources as resources

# options

allows[output] {
    resource := resources[_]
    lgtemapezqleqyhyzryw
    output := {"id": resource.id, "type": resource.type, "allow": true}
}

allows[output] {
    resource := resources[_]
    jjpjzpfrfegmotafeths(resource)
    output := {"id": resource.id, "type": resource.type, "allow": true}
}

# rules

lgtemapezqleqyhyzryw {
    "admin" in user.roles
}

jjpjzpfrfegmotafeths(resource) {
    resource.attributes.owner_id == "1"
}
```


## Need More, Check Out our API

Permify API is an authorization API which you can add complex rbac and abac solutions.

[<img src="https://user-images.githubusercontent.com/39353278/157747851-ea8462be-60a4-498e-872a-e44cf42411b0.png" width="419px" />](https://www.permify.co/get-started)


<h2 align="left">:heart: Let's get connected:</h2>

<p align="left">
<a href="https://twitter.com/GetPermify">
  <img alt="guilyx | Twitter" width="50px" src="https://user-images.githubusercontent.com/43545812/144034996-602b144a-16e1-41cc-99e7-c6040b20dcaf.png"/>
</a>
<a href="https://www.linkedin.com/company/permifyco">
  <img alt="guilyx's LinkdeIN" width="50px" src="https://user-images.githubusercontent.com/43545812/144035037-0f415fc7-9f96-4517-a370-ccc6e78a714b.png" />
</a>
<a href="https://discord.gg/MJbUjwskdH">
  <img alt="guilyx's Discord" width="50px" src="https://www.apkmirror.com/wp-content/uploads/2021/06/09/60dbb1f8b30bb.png" />
</a>
</p>