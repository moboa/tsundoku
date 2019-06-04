package mangastream

import (
	"testing"

	"github.com/gocolly/colly"
)

func TestFetchPageImages(t *testing.T) {
	collector := colly.NewCollector()
	chapterURL := "https://readms.net/r/katekyo_hitman_reborn/409%20%28End%29/1593"
	numOfPages := 21

	images := FetchPageImages(collector, chapterURL)

	if len(images) != numOfPages {
		t.Errorf("Number of pages downloaded was incorrect, got: %d, want: %d.", len(images), numOfPages)
	}
}
