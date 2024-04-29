package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppError(t *testing.T) {
	err := NewError("error_code", Internal, High, "Error while running the test").SetDetail("More context about the error")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Internal, err.Type())
	assert.Equal(t, High, err.Severity())
	assert.Equal(t, "Error while running the test", err.Error())
	assert.Equal(t, "More context about the error", err.Detail())
	assert.NoError(t, err.Cause())
}

func TestAppCriticalErrorWrap(t *testing.T) {
	cause := errors.New("Inner Error that trigger the AppError")
	err := NewError("error_code", Internal, High, "Error while running the test").Wrap(cause).SetSeverity(Critical)

	assert.ErrorIs(t, cause, err.Cause())
	assert.Equal(t, Critical, err.Severity())
}

func TestNewValidationError(t *testing.T) {
	expectedValidationErrs := FieldValidationErrors{
		"name": []string{"name is empty"},
		"age":  []string{"user is under 18", "user must be an adult", "user is too young"},
	}
	err := NewValidationError("validate_user", "Error Validating User")
	err.AddValidationError(NewFieldValidationError("name", "name is empty"))
	err.AddValidationError(NewFieldValidationError("age", "user is under 18", "user must be an adult"))
	err.AddValidationError(NewFieldValidationError("age", "user is too young"))

	assert.Equal(t, "validate_user", err.Code())
	assert.Equal(t, Validation, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Validating User", err.Error())
	assert.NoError(t, err.Cause())
	assert.Equal(t, expectedValidationErrs, err.ValidationErrors())
}

func TestNewInternalError(t *testing.T) {
	err := NewInternalError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Internal, err.Type())
	assert.Equal(t, High, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, NotFound, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}

func TestNewPermissionError(t *testing.T) {
	err := NewPermissionError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Permission, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Unauthorized, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}

func TestNewConflictError(t *testing.T) {
	err := NewConflictError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Conflict, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}

func TestNewTimeoutError(t *testing.T) {
	err := NewTimeoutError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Timeout, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}

func TestNewCancelledError(t *testing.T) {
	err := NewCancelledError("error_code", "Error Message")

	assert.Equal(t, "error_code", err.Code())
	assert.Equal(t, Cancelled, err.Type())
	assert.Equal(t, Low, err.Severity())
	assert.Equal(t, "Error Message", err.Error())
	assert.NoError(t, err.Cause())
}
