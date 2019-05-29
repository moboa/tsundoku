package mangareader

import (
	"net/url"
	"testing"

	"github.com/gocolly/colly"
)

func TestFetchPageImages(t *testing.T) {
	collector := colly.NewCollector()
	chapterURL, _ := url.Parse("https://www.mangareader.net/yu-yu-hakusho/2")
	numOfPages := 24

	images := FetchPageImages(collector, chapterURL)

	if len(images) != numOfPages {
		t.Errorf("Number of pages downloaded was incorrect, got: %d, want: %d.", len(images), numOfPages)
	}
}
