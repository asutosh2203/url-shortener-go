package utils

import (
	"net/url"
	"regexp"
)

func IsValidURL(input string) bool {
	// Use regex to check the basic structure of the URL (optional)
	re := regexp.MustCompile(`^(http|https)://`)
	if !re.MatchString(input) {
		return false
	}

	// Parse the URL using net/url package
	parsedURL, err := url.Parse(input)
	if err != nil {
		return false
	}

	// Check if the URL has a valid scheme (http or https)
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Check if the URL has a valid host (domain or IP address)
	if parsedURL.Host == "" {
		return false
	}

	// If all checks pass, the URL is considered valid
	return true
}
