package app

// ErrorType defines the type of an error
type ErrorType string

const (
	// Internal error
	Internal ErrorType = "internal"
	// NotFound error
	NotFound ErrorType = "not_found"
	// Validation error
	Validation ErrorType = "validation"
	// Permission error
	Permission ErrorType = "permission"
	// Unauthorized error
	Unauthorized ErrorType = "unauthorized"
	// Conflict error
	Conflict ErrorType = "conflict"
	// Timeout error
	Timeout ErrorType = "timeout"
	// Cancelled error
	Cancelled ErrorType = "cancelled"
)

// Error Severity
type ErrorSeverity string

const (
	// Low - expected errors like validation
	Low ErrorSeverity = "low"
	// Medium - errors that are not urgent to look into.
	Medium ErrorSeverity = "medium"
	// High - errors that shouldn't happen but the system is still working.
	High ErrorSeverity = "high"
	// Critical - errors that impact the system from working.
	Critical ErrorSeverity = "critical"
)

// Validation Error
type FieldValidationErrors map[string][]string

type FieldValidationError struct {
	identifier string
	messages   []string
}

// New Validation Error
func NewFieldValidationError(identifier string, messages ...string) FieldValidationError {
	return FieldValidationError{
		identifier: identifier,
		messages:   messages,
	}
}

// Application Error
type Error struct {
	code             string
	errType          ErrorType
	severity         ErrorSeverity
	message          string
	detail           string
	cause            error
	validationErrors FieldValidationErrors
}

// Error message
func (e Error) Error() string {
	return e.message
}

// Error message
func (e Error) Code() string {
	return e.code
}

// Error message
func (e Error) Type() ErrorType {
	return e.errType
}

// Get the original error cause
func (e Error) Cause() error {
	return e.cause
}

// Get the original error cause
func (e Error) Severity() ErrorSeverity {
	return e.severity
}

// Set the original error severity
func (e Error) SetSeverity(severity ErrorSeverity) Error {
	e.severity = severity
	return e
}

// Get the error detail
func (e Error) Detail() string {
	return e.detail
}

// Set the erro detail
func (e Error) SetDetail(detail string) Error {
	e.detail = detail
	return e
}

// wrap the original error cause
func (e Error) Wrap(err error) Error {
	e.cause = err
	return e
}

// Get validation errors
func (e Error) ValidationErrors() FieldValidationErrors {
	return e.validationErrors
}

// Add validation errors
func (e *Error) AddValidationError(err FieldValidationError) {
	if e.validationErrors == nil {
		e.validationErrors = make(map[string][]string)
	}

	if value, ok := e.validationErrors[err.identifier]; ok {
		e.validationErrors[err.identifier] = append(value, err.messages...)
	} else {
		e.validationErrors[err.identifier] = err.messages
	}
}

// NewError creates a application new error
func NewError(code string, errType ErrorType, severity ErrorSeverity, msg string) Error {
	return Error{
		code:     code,
		errType:  errType,
		severity: severity,
		message:  msg,
	}
}

// NewInternalError creates an error of type internal
func NewInternalError(code string, msg string) Error {
	return NewError(code, Internal, High, msg)
}

// NewValidationError creates an error of type validation
func NewValidationError(code string, msg string) Error {
	return NewError(code, Validation, Low, msg)
}

// NewNotFoundError creates an error of type not found
func NewNotFoundError(code string, msg string) Error {
	return NewError(code, NotFound, Low, msg)
}

// NewPermissionError creates an error of type permission
func NewPermissionError(code string, msg string) Error {
	return NewError(code, Permission, Low, msg)
}

// NewUnauthorizedError creates an error of type unauthorized
func NewUnauthorizedError(code string, msg string) Error {
	return NewError(code, Unauthorized, Low, msg)
}

// NewConflictError creates an error of type conflict
func NewConflictError(code string, msg string) Error {
	return NewError(code, Conflict, Low, msg)
}

// NewTimeoutError creates an error of type timeout
func NewTimeoutError(code string, msg string) Error {
	return NewError(code, Timeout, Low, msg)
}

// NewCancelledError creates an error of type cancelled
func NewCancelledError(code string, msg string) Error {
	return NewError(code, Cancelled, Low, msg)
}
