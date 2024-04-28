
[![codecov](https://codecov.io/gh/Talento90/goliath/graph/badge.svg?token=4AIPK4UXUO)](https://codecov.io/gh/Talento90/goliath)
![build](https://github.com/Talento90/goliath/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Talento90/goliath)](https://goreportcard.com/report/github.com/Talento90/goliath)

<p align="center">
    
</p>
<p align="center">
    <img src="./assets/logo.png" alt="logo" width="200" >
</p>

<p align="center">
    <img src="./assets/goliath.webp" alt="goliath" width="250" >
</p>

# 📝 Summary

Goliath is an opinionated set of libraries to build resiliant, scalable and maintable applications. The goal of these libraries is to allow focusing on building applications instead of reinventing the wheel.

# 🚀 Features

- [apperror](/apperror) - create elegant application errors
- [appcontext](/appcontext/) - wrapper around the native `ctx.Context`
- [retry](/retry/) - retry a specific task securely
- [clock](/clock) - wrapper around `time.Now` to help during testing
- [sleep](/sleep) - wrapper around `time.Sleep` for testing

# 👀 Examples

### apperror
```go
// create application errors
err := NewValidation("validate_user", "Error Validating User")
err.AddValidationError(NewValidationError("name", "name is empty"))
err.AddValidationError(NewValidationError("age", "user is under 18", "user must be an adult"))

// wrap the inner cause of the error
conn, err := db.Connect(...)

if err != nil {
   return NewInternal("database_connection", "Error connecting to the database").Wrap(err)
}
```

### appcontext
```go
func Hello(w http.ResponseWriter, r *http.Request) {
    ctx := appcontext.FromContext(r.Context())
    traceId := appCtx.TraceID()
    userID, checkUser := ctx.UserID()

    if checkUser {
        //request authorized
    }
}
```

### retry
```go
	var GetPersonTask = func() (Person, error) {
        var p Person
        resp, err := http.Get("http://example.com/")

        if err != nil {
            return p, err
        }

        err = json.NewDecoder(resp.Body).Decode(&p)

		return resp, err
	}

    config := retry.NewConfig(3)

	result, err := retry.Execute(config, task)
```

### clock
```go
	clock := NewUtcClock()
	timeNowUtc := clock.Now().Format(time.RFC822)
```

### sleep
```go
	sleeper := sleep.New()
	sleeper.Sleep(1000)
```





