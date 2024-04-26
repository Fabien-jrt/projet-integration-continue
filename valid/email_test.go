package valid

import (
	"strings"
	"testing"
)

func TestIsEmailValid(t *testing.T) {
	// Test cases for valid emails
	validEmails := []string{
		"test@example.com",
		"john.doe@example.com",
		"jane.doe@example.co.uk",
	}

	for _, email := range validEmails {
		if !IsEmailValid(email) {
			t.Errorf("Expected %s to be a valid email", email)
		}
	}

	// Test cases for invalid emails
	invalidEmails := []string{
		"test",
		"test@example",
		"test@example.",
		"test@example..com",
		"test@example_com",
	}

	for _, email := range invalidEmails {
		if IsEmailValid(email) {
			t.Errorf("Expected %s to be an invalid email", email)
		}
	}
}

func TestIsTooLong(t *testing.T) {
	// Test cases for emails that are too long
	longEmails := []string{
		"test" + strings.Repeat("a", maxTotalChar),
		"test@example.com" + strings.Repeat("a", maxTotalChar),
	}

	for _, email := range longEmails {
		if !isTooLong(email) {
			t.Errorf("Expected %s to be too long", email)
		}
	}

	// Test cases for emails that are not too long
	shortEmails := []string{
		"test",
		"test@example.com",
	}

	for _, email := range shortEmails {
		if isTooLong(email) {
			t.Errorf("Expected %s to not be too long", email)
		}
	}
}
