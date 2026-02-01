package utils

import "errors"

// Common application errors
var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
)

// Custom error type with code
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Common error constructors

// NewValidationError creates a validation error
func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    "VALIDATION_ERROR",
		Message: message,
	}
}

// NewAuthError creates an authentication error
func NewAuthError(message string) *AppError {
	return &AppError{
		Code:    "AUTH_ERROR",
		Message: message,
	}
}

// NewDatabaseError creates a database error
func NewDatabaseError(message string) *AppError {
	return &AppError{
		Code:    "DATABASE_ERROR",
		Message: message,
	}
}

// NewInternalServer creates an internal server error
func NewInternalServer(message string) *AppError {
	return &AppError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: message,
	}
}

// NewNotFound creates a not found error
func NewNotFound(message string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: message,
	}
}

// NewConflictError creates a conflict error
func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    "CONFLICT",
		Message: message,
	}
}

// NewUnauthorizedError creates an unauthorized error
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    "UNAUTHORIZED",
		Message: message,
	}
}

// NewForbiddenError creates a forbidden error
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:    "FORBIDDEN",
		Message: message,
	}
}
