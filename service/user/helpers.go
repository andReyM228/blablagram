package user

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// IsEmailValid checks if email is valid
func IsEmailValid(email string) bool {
	regex := `^[a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`
	valid := regexp.MustCompile(regex).MatchString(email)

	return valid
}

// PasswordValidation checks if password is valid
func PasswordValidation(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return fmt.Errorf("the password is too short (%d characters), must be at least 8 characters", len(password))
	}

	// Checking for the presence of at least one letter
	hasLetter := regexp.MustCompile(`[a-zA-Z]`)
	if !hasLetter.MatchString(password) {
		return fmt.Errorf("the password must include at least one letter")
	}

	// Checking for the presence of at least one digit
	hasNumber := regexp.MustCompile(`\d`)
	if !hasNumber.MatchString(password) {
		return fmt.Errorf("the password must include at least one number")
	}

	return nil
}

func hashPasswordWithSalt(password string, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt + password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
