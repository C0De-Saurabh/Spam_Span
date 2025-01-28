package utils

import (
	"regexp"
)

// IsValidPhone validates the phone number format.
func IsValidPhone(phone string) bool {
	// Regular expression for validating a phone number
	// Assumes a phone number with 10 digits (modify this regex to support your country format)
	phoneRegex := `^(\+?[0-9]{1,4}?)?(\(?[0-9]{1,3}\)?[-.\s]?)?([0-9]{3})[-.\s]?([0-9]{3})[-.\s]?([0-9]{4})$`

	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}

func IsValidEmail(email string) bool {
	// Regular expression to match a basic email pattern
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// IsValidName validates the name format using regex
func IsValidName(name string) bool {
	// Validate name format using regex: only letters and spaces allowed
	nameRegex := `^[a-zA-Z\s]+$`
	return regexp.MustCompile(nameRegex).MatchString(name)
}
