package app

// ErrorType defines the type of an error
type ErrorType string

const (
	// ErrorInternal error
	ErrorInternal ErrorType = "internal"
	// ErrorNotFound error
	ErrorNotFound ErrorType = "not_found"
	// ErrorValidation error
	ErrorValidation ErrorType = "validation"
	// ErrorPermission error
	ErrorPermission ErrorType = "permission"
	// ErrorUnauthorised error
	ErrorUnauthorised ErrorType = "unauthorised"
	// ErrorConflict error
	ErrorConflict ErrorType = "conflict"
	// ErrorTimeout error
	ErrorTimeout ErrorType = "timeout"
	// ErrorCancelled error
	ErrorCancelled ErrorType = "cancelled"
)

// Error Severity
type ErrorSeverity string

const (
	// ErrorSeverityLow - expected errors like validation
	ErrorSeverityLow ErrorSeverity = "low"
	// ErrorSeverityMedium - errors that are not urgent to look into.
	ErrorSeverityMedium ErrorSeverity = "medium"
	// ErrorSeverityHigh - errors that shouldn't happen but the system is still working.
	ErrorSeverityHigh ErrorSeverity = "high"
	// ErrorSeverityCritical - errors that impact the system from working.
	ErrorSeverityCritical ErrorSeverity = "critical"
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

// Error code
func (e Error) Code() string {
	return e.code
}

// Error type
func (e Error) Type() ErrorType {
	return e.errType
}

// Get the original error cause
func (e Error) Cause() error {
	return e.cause
}

// Get the error severity
func (e Error) Severity() ErrorSeverity {
	return e.severity
}

// Set the original error severity
func (e *Error) SetSeverity(severity ErrorSeverity) *Error {
	e.severity = severity
	return e
}

// Get the error detail
func (e Error) Detail() string {
	return e.detail
}

// Set the error detail
func (e *Error) SetDetail(detail string) *Error {
	e.detail = detail
	return e
}

// wrap the original error cause
func (e *Error) Wrap(err error) *Error {
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
func NewError(code string, errType ErrorType, severity ErrorSeverity, msg string) *Error {
	return &Error{
		code:     code,
		errType:  errType,
		severity: severity,
		message:  msg,
	}
}

// NewErrorInternal creates an error of type internal
func NewErrorInternal(code string, msg string) *Error {
	return NewError(code, ErrorInternal, ErrorSeverityHigh, msg)
}

// NewErrorValidation creates an error of type validation
func NewErrorValidation(code string, msg string) *Error {
	return NewError(code, ErrorValidation, ErrorSeverityLow, msg)
}

// NewErrorNotFound creates an error of type not found
func NewErrorNotFound(code string, msg string) *Error {
	return NewError(code, ErrorNotFound, ErrorSeverityLow, msg)
}

// NewErrorPermission creates an error of type permission
func NewErrorPermission(code string, msg string) *Error {
	return NewError(code, ErrorPermission, ErrorSeverityLow, msg)
}

// NewErrorUnauthorised creates an error of type unauthorised
func NewErrorUnauthorised(code string, msg string) *Error {
	return NewError(code, ErrorUnauthorised, ErrorSeverityLow, msg)
}

// NewErrorConflict creates an error of type conflict
func NewErrorConflict(code string, msg string) *Error {
	return NewError(code, ErrorConflict, ErrorSeverityLow, msg)
}

// NewErrorTimeout creates an error of type timeout
func NewErrorTimeout(code string, msg string) *Error {
	return NewError(code, ErrorTimeout, ErrorSeverityLow, msg)
}

// NewErrorCancelled creates an error of type cancelled
func NewErrorCancelled(code string, msg string) *Error {
	return NewError(code, ErrorCancelled, ErrorSeverityLow, msg)
}
