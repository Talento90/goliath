
[![codecov](https://codecov.io/gh/Talento90/goliath/graph/badge.svg?token=4AIPK4UXUO)](https://codecov.io/gh/Talento90/goliath)
![build](https://github.com/Talento90/goliath/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Talento90/goliath)](https://goreportcard.com/report/github.com/Talento90/goliath)
[![GoDoc](https://godoc.org/github.com/Talento90/goliath?status.svg)](https://godoc.org/github.com/Talento90/goliath)

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
- [httperror](/httperror) - implementation of the [RFC7807 Problem Details](https://datatracker.ietf.org/doc/html/rfc7807)

# 👀 Examples

### apperror
```go
// create application errors
err := apperror.NewValidation("validate_user", "Error Validating User")
err.AddValidationError(NewValidationError("name", "name is empty"))
err.AddValidationError(NewValidationError("age", "user is under 18", "user must be an adult"))

// wrap the inner cause of the error
conn, err := db.Connect(...)

if err != nil {
   return apperror.NewInternal("database_connection", "Error connecting to the database").SetSeverity(apperror.Critical).Wrap(err)
}

// create a raw error
err := apperror.New(("error_code", apperror.Internal, apperror.High, "Error message"))
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


### httperror
```go
	appCtx := appcontext.FromContext(ctx)
	err := apperror.NewValidation("invalid_payment_data", "The payment request is invalid")
	err.AddValidationError(apperror.NewValidationError("amount", "Amount needs to be positive"))
	err.AddValidationError(apperror.NewValidationError("currency", "currency is required"))

	httpErr := New(appCtx, err, "/payments")
```
*Problem Detail Output*
```json
{
  "type": "invalid_payment_data",
  "title": "The payment request is invalid",
  "status": 400,
  "instance": "/payments",
  "traceId": "9b1b4579-b455-4eed-ac80-923668593dcc",
  "errors": {
    "amount": ["Amount needs to be positive"],
    "currency": ["currency is required"]
  }
}
```




