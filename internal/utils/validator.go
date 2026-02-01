package utils

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrPasswordTooShort = errors.New("password must be at least 6 characters")
	ErrEmptyField       = errors.New("field cannot be empty")
)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return ErrEmptyField
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return ErrPasswordTooShort
	}
	return nil
}

// ValidateName validates name field
func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyField
	}
	return nil
}

// SanitizeString removes leading/trailing whitespace
func SanitizeString(s string) string {
	return strings.TrimSpace(s)
}
