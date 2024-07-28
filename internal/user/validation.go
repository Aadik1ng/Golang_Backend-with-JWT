package user

import (
	"errors"
	"regexp"
)

func ValidateUser(user User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	// Simple email regex validation
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email format")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Mobile == "" {
		return errors.New("mobile number is required")
	}

	// Simple mobile number regex validation (assuming 10-digit numbers)
	mobileRegex := regexp.MustCompile(`^\d{10}$`)
	if !mobileRegex.MatchString(user.Mobile) {
		return errors.New("invalid mobile number format")
	}

	return nil
}
