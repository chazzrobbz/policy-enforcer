
<h1 align="center">
  <img alt="policy enforcer logo" src="https://user-images.githubusercontent.com/39353278/165170342-d1d61da0-f464-48c5-a3df-6a9bf1d14aa2.png" width="250px"/><br/>
  Policy Enforcer
</h1>

<p align="center">Represent your rego rules programmatically.</p>

<p align="center">Rego is the policy language for defining rules that are evaluated by the OPA (Open Policy Agent) engine.</p>


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

## 🚲 Basic Usage

```go
var user = User{
    ID:   1,
    Roles:  []string{"admin"},
    Attributes: map[string]interface{}{
            "tenure": 8,
    },
}

var isAdmin = enforcer.NewRule("'admin' in user.roles").SetFailMessage("user is not an admin").SetKey("is admin")
var isSenior = enforcer.NewRule("user.tenure > 8").SetFailMessage("user is not senior")
var isManager = enforcer.NewRule("'manager' in user.roles").SetFailMessage("user is not manager")

enforcer.New()
    
// set user object
enforcer.Set("user", user)

// its means the user must be either an admin or a senior manager
enforcer.Option(isAdmin).Option(isSenior, isManager)
   
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
            Allow: false, // is manager result
            Key: "xoeffrswxpldnjobcsnv", // if the key is not set it is created automatically
            Message: "user is not manager"
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