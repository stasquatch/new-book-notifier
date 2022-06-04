package main

import (
	"time"

	"github.com/stasquatch/new-book-notifier/internal"
)

var (
	// consider any new books to be published within the last week
	STARTING_TIMEFRAME = time.Now().Add(-7 * 24 * time.Hour)
	authorList         = []string{
		"Katie Ruggle",
	}
)

func main() {
	// kick off the whole shebang!
	internal.ProcessBookData(authorList, STARTING_TIMEFRAME)
}
