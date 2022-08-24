package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Add NonFieldErrors for errors not related to a specific form field
// Define a new Validator type which contains a map of validation errors
type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

// Use regexp.MustCompile() to check format of email address
// Returns a pointer to a compiled regexp.Regexp type
// Or panics in the event of an error
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Valid() returns true if the FieldErrors map is empty
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

// AddFieldError() adds an error message to the FieldErrors map
func (v *Validator) AddFieldError(key, message string) {
	// Init the map
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// Adds error messages to NonFieldErrors slice
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

// CheckField() adds an error message to the map if a validation check is not ok
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank() returns true if a value is not an empty string
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() returns true if value is less than n chars
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt() returns true if the value is in the list of ints
func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

// Returns true if a value contains n characters
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Returns true if a value matches the regular expression
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
