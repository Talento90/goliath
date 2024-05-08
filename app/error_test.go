package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAppError(t *testing.T) {
	err := NewError("error_code", Internal, High, "Error while running the test").SetDetail("More context about the error")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Internal, err.Type())
	require.Equal(t, High, err.Severity())
	require.Equal(t, "Error while running the test", err.Error())
	require.Equal(t, "More context about the error", err.Detail())
	require.NoError(t, err.Cause())
}

func TestAppCriticalErrorWrap(t *testing.T) {
	cause := errors.New("Inner Error that trigger the AppError")
	err := NewError("error_code", Internal, High, "Error while running the test").Wrap(cause).SetSeverity(Critical)

	require.ErrorIs(t, cause, err.Cause())
	require.Equal(t, Critical, err.Severity())
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

	require.Equal(t, "validate_user", err.Code())
	require.Equal(t, Validation, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Validating User", err.Error())
	require.NoError(t, err.Cause())
	require.Equal(t, expectedValidationErrs, err.ValidationErrors())
}

func TestNewInternalError(t *testing.T) {
	err := NewInternalError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Internal, err.Type())
	require.Equal(t, High, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, NotFound, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewPermissionError(t *testing.T) {
	err := NewPermissionError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Permission, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewUnauthorisedError(t *testing.T) {
	err := NewUnauthorisedError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Unauthorised, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewConflictError(t *testing.T) {
	err := NewConflictError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Conflict, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewTimeoutError(t *testing.T) {
	err := NewTimeoutError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Timeout, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewCancelledError(t *testing.T) {
	err := NewCancelledError("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, Cancelled, err.Type())
	require.Equal(t, Low, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}
