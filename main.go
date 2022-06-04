package main

import (
	"log"
	"time"

	"github.com/stasquatch/new-book-notifier/internal"
	"github.com/stasquatch/new-book-notifier/internal/helpers"
)

// consider any new books to be published within the last week
var STARTING_TIMEFRAME = time.Now().Add(-7 * 24 * time.Hour)

func main() {
	authorList := []string{
		"Katie Ruggle",
	}

	for i := range authorList {
		authorName := authorList[i]

		data, err := internal.GetBooksByAuthor(authorName)
		if err != nil {
			log.Panicln(err)
		}

		newBookTitles := make([]string, 0)
		for i := range data {
			book := data[i].VolumeInfo

			// if they aren't the author, skip to next book
			if isAuthor := helpers.IsAuthorOfBook(authorName, book.Authors); !isAuthor {
				continue
			}

			isNew, err := helpers.IsNewBook(book.PublishedDate, STARTING_TIMEFRAME)
			if err != nil {
				log.Printf("error with book %s, skipping. error: %s", book.Title, err.Error())
				continue
			}
			if isNew {
				newBookTitles = append(newBookTitles, book.Title)
			}
		}

		if len(newBookTitles) == 0 {
			log.Printf("no new books for author %s", authorName)
		} else {
			log.Printf("at least one new book for author %s", authorName)
		}
	}
}
