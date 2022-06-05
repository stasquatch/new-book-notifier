package helpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsAuthorOfBook(t *testing.T) {
	arr := []string{"one", "two", "three"}
	cases := []struct {
		Name       string
		AuthorName string
		Arr        []string
		Expected   bool
	}{
		{"returns true if author is in array", "one", arr, true},
		{"returns false if author is not in array", "five", arr, false},
		{"returns false if array is empty", "one", []string{}, false},
		{"returns true if author matches case insensitive", "ONE", arr, true},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Expected, IsAuthorOfBook(c.AuthorName, c.Arr))
		})
	}
}

func TestIsNewBook(t *testing.T) {
	now := time.Now()
	oneWeekAgo := now.Add(-7 * 24 * time.Hour)
	yesterday := now.Add(-24 * time.Hour)
	twoWeeksAgo := now.Add(-7 * 24 * 2 * time.Hour)
	cases := []struct {
		Name       string
		Date       string
		Expected   bool
		ErrMessage *string
	}{
		{"returns true if date is after starting timeframe", yesterday.Format("2006-01-02"), true, nil},
		{"returns false if date is before starting timeframe", twoWeeksAgo.Format("2006-01-02"), false, nil},
		{"returns error for parsing error", "wrong", false, strPtr("parsing time \"wrong\" as \"2006-01-02\": cannot parse \"wrong\" as \"2006\"")},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result, err := IsNewBook(c.Date, oneWeekAgo)
			assert.Equal(t, c.Expected, result)
			if err != nil {
				assert.Equal(t, *c.ErrMessage, err.Error())
			}
		})
	}
}

func TestFormatNewBooksAsString(t *testing.T) {
	t.Run("successfully concats a map of authors/books", func(t *testing.T) {
		b := make(map[string][]string, 2)
		b["one"] = []string{"single book title"}
		b["two"] = []string{"two books", "will go here"}
		result := FormatNewBooksAsString(b)
		assert.Equal(t, `The following authors have published something new!

one: single book title

two: two books, will go here`, result)
	})

	t.Run("successfully concats a map of a single author/book", func(t *testing.T) {
		b := make(map[string][]string, 1)
		b["one"] = []string{"single"}
		result := FormatNewBooksAsString(b)
		assert.Equal(t, `The following author has published something new!

one: single`, result)
	})

	t.Run("returns empty string if map is empty", func(t *testing.T) {
		b := make(map[string][]string, 0)
		result := FormatNewBooksAsString(b)
		assert.Empty(t, result)
	})
}
