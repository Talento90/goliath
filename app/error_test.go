package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAppError(t *testing.T) {
	err := NewError("error_code", ErrorInternal, ErrorSeverityHigh, "Error while running the test").SetDetail("More context about the error")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorInternal, err.Type())
	require.Equal(t, ErrorSeverityHigh, err.Severity())
	require.Equal(t, "Error while running the test", err.Error())
	require.Equal(t, "More context about the error", err.Detail())
	require.NoError(t, err.Cause())
}

func TestAppCriticalErrorWrap(t *testing.T) {
	cause := errors.New("Inner Error that trigger the AppError")
	err := NewError("error_code", ErrorInternal, ErrorSeverityHigh, "Error while running the test").Wrap(cause).SetSeverity(ErrorSeverityCritical)

	require.ErrorIs(t, cause, err.Cause())
	require.Equal(t, ErrorSeverityCritical, err.Severity())
}

func TestNewValidationError(t *testing.T) {
	expectedValidationErrs := FieldValidationErrors{
		"name": []string{"name is empty"},
		"age":  []string{"user is under 18", "user must be an adult", "user is too young"},
	}
	err := NewErrorValidation("validate_user", "Error Validating User")
	err.AddValidationError(NewFieldValidationError("name", "name is empty"))
	err.AddValidationError(NewFieldValidationError("age", "user is under 18", "user must be an adult"))
	err.AddValidationError(NewFieldValidationError("age", "user is too young"))

	require.Equal(t, "validate_user", err.Code())
	require.Equal(t, ErrorValidation, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Validating User", err.Error())
	require.NoError(t, err.Cause())
	require.Equal(t, expectedValidationErrs, err.ValidationErrors())
}

func TestNewInternalError(t *testing.T) {
	err := NewErrorInternal("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorInternal, err.Type())
	require.Equal(t, ErrorSeverityHigh, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewErrorNotFound("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorNotFound, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewPermissionError(t *testing.T) {
	err := NewErrorPermission("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorPermission, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewUnauthorisedError(t *testing.T) {
	err := NewErrorUnauthorised("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorUnauthorised, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewConflictError(t *testing.T) {
	err := NewErrorConflict("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorConflict, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewTimeoutError(t *testing.T) {
	err := NewErrorTimeout("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorTimeout, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}

func TestNewCancelledError(t *testing.T) {
	err := NewErrorCancelled("error_code", "Error Message")

	require.Equal(t, "error_code", err.Code())
	require.Equal(t, ErrorCancelled, err.Type())
	require.Equal(t, ErrorSeverityLow, err.Severity())
	require.Equal(t, "Error Message", err.Error())
	require.NoError(t, err.Cause())
}
