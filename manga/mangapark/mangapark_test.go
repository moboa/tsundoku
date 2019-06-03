package mangapark

import (
	"net/url"
	"testing"

	"github.com/gocolly/colly"
)

func TestFetchPageImages(t *testing.T) {
	collector := colly.NewCollector()
	chapterURL, _ := url.Parse("https://mangapark.net/manga/ranma-1-2-rumiko-takahashi/i1347243/")
	numOfPages := 20

	images := FetchPageImages(collector, chapterURL)

	if len(images) != numOfPages {
		t.Errorf("Number of pages downloaded was incorrect, got: %d, want: %d.", len(images), numOfPages)
	}
}
