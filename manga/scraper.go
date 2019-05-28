package manga

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/moboa/tsundoku/manga/mangareader"
)

var collector = colly.NewCollector()

var sourceParsers = map[string]func(*colly.Collector, *url.URL) []string{
	"www.mangareader.net": mangareader.FetchPageImages,
}

// FetchPageImages returns a list containing images of the chapter at specified URL
func FetchPageImages(chapterURL *url.URL) []string {
	sourceParser := sourceParsers[chapterURL.Hostname()]

	if sourceParser == nil {
		log.Fatal(chapterURL.Hostname() + " is not supported.")
	}

	images := sourceParser(collector, chapterURL)
	lenImages := len(images)
	fmt.Println("Downloaded " + strconv.Itoa(lenImages) + " pages from " + chapterURL.String())

	return images
}
