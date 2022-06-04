package helpers

import (
	"strings"
	"time"
)

func IsAuthorOfBook(authorName string, authorsArr []string) bool {
	for i := range authorsArr {
		if strings.EqualFold(authorsArr[i], authorName) {
			return true
		}
	}
	return false
}

func IsNewBook(publishedDate string, startingTimeframe time.Time) (bool, error) {
	pubDate, err := time.Parse("2006-01-02", publishedDate)
	if err != nil {
		return false, err
	}

	return pubDate.After(startingTimeframe), nil
}

func strPtr(s string) *string {
	return &s
}
