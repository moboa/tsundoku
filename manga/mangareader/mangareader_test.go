package mangareader

import (
	"testing"

	"github.com/gocolly/colly"
)

var collector = colly.NewCollector()

func TestFetchMangareaderPageImages(t *testing.T) {
	chapterURL := "https://www.mangareader.net/yu-yu-hakusho/2"
	numOfPages := 24

	images := FetchPageImages(collector, chapterURL)

	if len(images) != numOfPages {
		t.Errorf("Number of pages downloaded was incorrect, got: %d, want: %d.", len(images), numOfPages)
	}
}

func TestFetchMangapandaPageImages(t *testing.T) {
	chapterURL := "https://www.mangapanda.com/noblesse/2"
	numOfPages := 22

	images := FetchPageImages(collector, chapterURL)

	if len(images) != numOfPages {
		t.Errorf("Number of pages downloaded was incorrect, got: %d, want: %d.", len(images), numOfPages)
	}
}
