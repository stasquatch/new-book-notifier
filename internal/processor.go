package internal

import (
	"errors"
	"time"

	"github.com/stasquatch/new-book-notifier/internal/helpers"
)

func ProcessBookData(authorList []string, startingTimeframe time.Time) (string, error) {
	allNewBooksByAuthor := make(map[string][]string)
	for i := range authorList {
		authorName := authorList[i]

		data, err := getBooksByAuthor(authorName)
		if err != nil {
			return "", errors.New(err.Error())
		}

		newBooks := filterNewBooksForAuthor(data, authorName, startingTimeframe)
		if len(newBooks) > 0 {
			allNewBooksByAuthor[authorName] = newBooks
		}
	}

	return helpers.FormatNewBooksAsString(allNewBooksByAuthor), nil
}

func filterNewBooksForAuthor(data []GoogleBookItem, authorName string, startingTimeframe time.Time) []string {
	newBookTitles := make([]string, 0)
	for i := range data {
		book := data[i].VolumeInfo

		// if they aren't the author, skip to next book
		if isAuthor := helpers.IsAuthorOfBook(authorName, book.Authors); !isAuthor {
			continue
		}

		isNew, err := helpers.IsNewBook(book.PublishedDate, startingTimeframe)
		if err != nil {
			// silently fail and continue
			continue
		}
		if isNew {
			newBookTitles = append(newBookTitles, book.Title)
		}
	}

	return newBookTitles
}
