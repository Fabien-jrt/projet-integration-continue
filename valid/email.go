package valid

import "strings"

func IsEmailValid(email string) bool {
	return false
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
	validExtensions := []string{"com", "net", "org"} // add more valid extensions if needed
	for _, ext := range validExtensions {
		if strings.ToLower(extension) == ext {
			return true
		}
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

	validDomains := []string{"gmail", "yahoo", "hotmail"} // add more valid domains if needed
	for _, dom := range validDomains {
		if strings.ToLower(domain) == dom {
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
