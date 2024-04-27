package apperror

// Type defines the type of an error
type Type string

const (
	// Internal error
	Internal Type = "internal"
	// NotFound error
	NotFound Type = "not_found"
	// Validation error
	Validation Type = "validation"
	// Permission error
	Permission Type = "permission"
	// Unauthorized error
	Unauthorized Type = "unauthorized"
	// Conflict error
	Conflict Type = "conflict"
	// Timeout error
	Timeout Type = "timeout"
	// Cancelled error
	Cancelled Type = "cancelled"
)

// Error Severity
type Severity string

const (
	// Internal error
	Low Severity = "low"
	// NotFound error
	Medium Severity = "medium"
	// Validation error
	High Severity = "high"
	// Timeout error
	Critical Severity = "critical"
)

// Validation Error
type ValidationErrors map[string][]string

type ValidationError struct {
	identifier string
	messages   []string
}

// New Validation Error
func NewValidationError(identifier string, messages ...string) ValidationError {
	return ValidationError{
		identifier: identifier,
		messages:   messages,
	}
}

// Application Error
type AppError struct {
	code             string
	errType          Type
	severity         Severity
	message          string
	cause            error
	validationErrors ValidationErrors
}

// Error message
func (e AppError) Error() string {
	return e.message
}

// Error message
func (e AppError) Code() string {
	return e.code
}

// Error message
func (e AppError) Type() Type {
	return e.errType
}

// Get the original error cause
func (e AppError) Cause() error {
	return e.cause
}

// Get the original error cause
func (e AppError) Severity() Severity {
	return e.severity
}

// Set the original error severity
func (e AppError) SetSeverity(severity Severity) AppError {
	e.severity = severity
	return e
}

// wrap the original error cause
func (e AppError) Wrap(err error) AppError {
	e.cause = err
	return e
}

// Get validation errors
func (e AppError) ValidationErrors() ValidationErrors {
	return e.validationErrors
}

// Add validation errors
func (e *AppError) AddValidationError(err ValidationError) {
	if e.validationErrors == nil {
		e.validationErrors = make(map[string][]string)
	}

	if value, ok := e.validationErrors[err.identifier]; ok {
		e.validationErrors[err.identifier] = append(value, err.messages...)
	} else {
		e.validationErrors[err.identifier] = err.messages
	}
}

// New creates a application new error
func New(code string, errType Type, severity Severity, msg string) AppError {
	return AppError{
		code:     code,
		errType:  errType,
		severity: severity,
		message:  msg,
	}
}

// NewInternal creates an error of type internal
func NewInternal(code string, msg string) AppError {
	return New(code, Internal, High, msg)
}

// NewValidation creates an error of type validation
func NewValidation(code string, msg string) AppError {
	return New(code, Validation, Low, msg)
}

// NewNotFound creates an error of type not found
func NewNotFound(code string, msg string) AppError {
	return New(code, NotFound, Low, msg)
}

// NewPermission creates an error of type permission
func NewPermission(code string, msg string) AppError {
	return New(code, Permission, Low, msg)
}

// NewUnauthorized creates an error of type unauthorized
func NewUnauthorized(code string, msg string) AppError {
	return New(code, Unauthorized, Low, msg)
}

// NewConflict creates an error of type conflict
func NewConflict(code string, msg string) AppError {
	return New(code, Conflict, Low, msg)
}

// NewTimeout creates an error of type timeout
func NewTimeout(code string, msg string) AppError {
	return New(code, Timeout, Low, msg)
}

// NewCancelled creates an error of type cancelled
func NewCancelled(code string, msg string) AppError {
	return New(code, Cancelled, Low, msg)
}
