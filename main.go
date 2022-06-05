package main

import (
	"fmt"
	"time"

	"github.com/stasquatch/new-book-notifier/internal"
)

var (
	// consider any new books to be published within the last week
	STARTING_TIMEFRAME = time.Now().Add(-7 * 24 * time.Hour)
	authorList         = []string{
		"Katie Ruggle",
		"Stephen Graham Jones",
		"Emily Henry",
		"Sally Thorne",
		"Abby Jimenez",
		"Talia Hibbert",
		"Emery Lee",
		"Annabeth Albert",
		"Jasmine Guillory",
		"Helen Hoang",
		"Leigh Bardugo",
		"Jaci Burton",
		"Mhairi McFarlane",
	}
)

func main() {
	// kick off the whole shebang!
	result, err := internal.ProcessBookData(authorList, STARTING_TIMEFRAME)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
