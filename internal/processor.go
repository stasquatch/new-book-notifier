package internal

import (
	"log"
	"time"

	"github.com/stasquatch/new-book-notifier/internal/helpers"
)

func ProcessBookData(authorList []string, startingTimeframe time.Time) {
	allNewBooksByAuthor := make(map[string][]string)
	for i := range authorList {
		authorName := authorList[i]

		data, err := getBooksByAuthor(authorName)
		if err != nil {
			log.Panicln(err)
		}

		newBooks := filterNewBooksForAuthor(data, authorName, startingTimeframe)
		if len(newBooks) > 0 {
			allNewBooksByAuthor[authorName] = newBooks
		}
	}
	// TODO: send email with contents of allNewBooksByAuthor
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
			log.Printf("error with book %s, skipping. error: %s", book.Title, err.Error())
			continue
		}
		if isNew {
			newBookTitles = append(newBookTitles, book.Title)
		}
	}

	return newBookTitles
}
