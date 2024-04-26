package valid

import (
	"io"
	"net/http"
	"strings"
)

const maxCharAfterAt = 255
const maxCharBeforeAt = 64
const maxTotalChar = maxCharAfterAt + maxCharBeforeAt
const allowedCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

func IsEmailValid(email string) bool {
	if isTooLong(email) {
		return false
	}
	return hasAtSign(email) &&
		hasValidExtention(email) &&
		hasValidDomain(email) &&
		hasValidLengthBeforeAt(email) &&
		hasValidLengthAfterAt(email) &&
		hasNoSpace(email) &&
		hasSomethingBeforeAt(email) &&
		hasSomethingAfterAt(email) &&
		hasSomethingAfterExt(email) &&
		!hasAdjacentDots(email)
}

func isTooLong(email string) bool {
	return len(email) > maxTotalChar
}

func hasAtSign(email string) bool {
	// check if the email has an @ sign
	for _, char := range email {
		if char == '@' {
			return true
		}
	}
	return false
}

func hasValidExtention(email string) bool {
	// check if the email has a valid extension
	parts := strings.Split(email, ".")
	extension := parts[len(parts)-1]

	// fetch the list of valid extensions from the URL
	resp, err := http.Get("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		// handle error
		return false
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return false
	}

	// convert the response body to a string
	extensions := strings.ToLower(string(body))

	// check if the extension is in the list of valid extensions
	if strings.Contains(extensions, strings.ToLower(extension)) {
		return true
	}

	return false
}

func hasValidDomain(email string) bool {
	// check if the email has a valid domain
	parts := strings.Split(email, "@")
	domain := parts[1]
	// get only the domain
	parts = strings.Split(domain, ".")
	domain = parts[len(parts)-1]

	// check for forbidden characters in the domain name
	for _, char := range domain {
		if !isValidDomainCharacter(char) {
			return false
		}
	}

	return true
}

func isValidDomainCharacter(char rune) bool {
	// check if the character is in the allowed set
	for _, allowedChar := range allowedCharacters {
		if char == allowedChar {
			return true
		}
	}

	return false
}

func hasValidLengthBeforeAt(email string) bool {
	// 64 char max before @
	parts := strings.Split(email, "@")
	beforeAt := parts[0]

	// Check if it's more than 64 characters
	if len(beforeAt) > 64 {
		return false
	}
	return true
}

func hasValidLengthAfterAt(email string) bool {
	// 255 char max after @
	parts := strings.Split(email, "@")
	afterAt := parts[1]

	// Check if it's more than 255 characters
	if len(afterAt) > 255 {
		return false
	}
	return true
}

func hasNoSpace(email string) bool {
	// check if the email has any spaces
	return !strings.Contains(email, " ")
}

func hasSomethingBeforeAt(email string) bool {
	// check if there is something before the @ sign
	parts := strings.Split(email, "@")
	beforeAt := parts[0]
	return len(beforeAt) > 0
}

func hasSomethingAfterAt(email string) bool {
	// check if there is something after the @ sign
	parts := strings.Split(email, "@")
	afterAt := parts[1]
	return len(afterAt) > 0
}

func hasSomethingAfterExt(email string) bool {
	// check if there is something after the . in the extension
	parts := strings.Split(email, ".")
	extension := parts[len(parts)-1]
	return len(extension) > 0
}

func hasAdjacentDots(email string) bool {
	// check if the email has adjacent dots
	for i := 0; i < len(email)-1; i++ {
		if email[i] == '.' && email[i+1] == '.' {
			return true
		}
	}
	return false
}
