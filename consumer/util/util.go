package util

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"strings"
	"time"
)

func HashItemName(input string) string {
	encodedString := url.PathEscape(input)

	//change %28 to ( and %29 to )
	encodedString = strings.ReplaceAll(encodedString, "%28", "(")
	encodedString = strings.ReplaceAll(encodedString, "%29", ")")
	encodedString = strings.ReplaceAll(encodedString, "%E2%98%85", "★")

	return encodedString
}

func ParseTimeAsUnix(input string) (int64, error) {
	unixTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		log.Error(err)
	}

	return unixTime.Unix(), err
}
