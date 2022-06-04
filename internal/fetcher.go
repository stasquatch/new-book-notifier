package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type GoogleBooksApiResponse struct {
	Kind       string           `json:"kind"`
	TotalItems int              `json:"totalItems"`
	Items      []GoogleBookItem `json:"items"`
}

type GoogleBookItem struct {
	Kind       string           `json:"kind"`
	Id         string           `json:"id"`
	VolumeInfo GoogleVolumeInfo `json:"volumeInfo"`
}

type GoogleVolumeInfo struct {
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	PublishedDate string   `json:"publishedDate"`
}

func getBooksByAuthor(authorName string) ([]GoogleBookItem, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&orderBy=newest", url.QueryEscape(authorName))
	resp, err := client.Get(url)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with request for %s: %s", authorName, err.Error()))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading request bdoy: %s", err.Error()))
	}

	var data GoogleBooksApiResponse

	err = json.Unmarshal([]byte(body), &data) // here!

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error unmarshalling google book response: %s", err.Error()))
	}

	return data.Items, nil
}
