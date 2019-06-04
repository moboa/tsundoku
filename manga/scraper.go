package manga

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/moboa/tsundoku/manga/mangapark"
	"github.com/moboa/tsundoku/manga/mangareader"
	"github.com/moboa/tsundoku/manga/mangastream"
)

var collector = colly.NewCollector()

var sourceParsers = map[string]func(*colly.Collector, string) []string{
	"mangareader": mangareader.FetchPageImages,
	"mangapanda":  mangareader.FetchPageImages, // mangapanda has the same layout as mangareader
	"mangapark":   mangapark.FetchPageImages,
	"readms":      mangastream.FetchPageImages,
}

// FetchPageImages returns a list containing images of the chapter at specified URL
func FetchPageImages(pageURL *url.URL) []string {
	r := regexp.MustCompile(`(www\.)?(.*)\..*`)
	key := r.FindStringSubmatch(pageURL.Hostname())[2]

	sourceParser := sourceParsers[key]

	if sourceParser == nil {
		log.Fatal(pageURL.Hostname() + " is not supported.")
	}

	r = regexp.MustCompile(`(.+)(/\d+)/?$`)
	chapterURL := r.FindStringSubmatch(pageURL.String())[1]
	images := sourceParser(collector, chapterURL)
	lenImages := len(images)
	fmt.Println("Downloaded " + strconv.Itoa(lenImages) + " pages from " + pageURL.String())

	return images
}
