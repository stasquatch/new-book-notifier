package helpers

import (
	"fmt"
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

func FormatNewBooksAsString(newBooks map[string][]string) string {
	if len(newBooks) == 0 {
		return ""
	}

	str := ""
	pluralAuthor := "author has"
	if len(newBooks) > 1 {
		pluralAuthor = "authors have"
	}
	str = fmt.Sprintf("The following %s published something new!", pluralAuthor)

	for author, bookTitles := range newBooks {
		if len(bookTitles) == 0 {
			continue
		}
		str += fmt.Sprintf("\n\n%s: %s", author, strings.Join(bookTitles, ", "))
	}

	return str
}

func strPtr(s string) *string {
	return &s
}
