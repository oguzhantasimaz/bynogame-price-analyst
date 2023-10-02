package pkg

import (
	"net/url"
	"strings"
)

func HashItemName(input string) string {
	encodedString := url.PathEscape(input)

	//change %28 to ( and %29 to )
	encodedString = strings.ReplaceAll(encodedString, "%28", "(")
	encodedString = strings.ReplaceAll(encodedString, "%29", ")")
	encodedString = strings.ReplaceAll(encodedString, "%E2%98%85", "★")

	return encodedString
}
