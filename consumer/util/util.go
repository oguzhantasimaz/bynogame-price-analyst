package util

import (
	"net/url"
	"strings"
)

func HashItemName(input string) string {
	encodedString := url.PathEscape(input)

	encodedString = strings.ReplaceAll(encodedString, "%E2%98%85", "★")

	return encodedString
}
