package mangapark

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/moboa/tsundoku/manga/mangautil"
)

// FetchPageImages returns a list containing images of the chapter at specified MangaPark URL
func FetchPageImages(collector *colly.Collector, chapterURL *url.URL) []string {
	imagesUrls := fetchImageUrls(chapterURL)
	return mangautil.FetchImages(collector, imagesUrls)
}

func fetchImageUrls(chapterURL *url.URL) []string {
	const splitAfterSep = "var _load_pages = "
	const splitBeforeSep = ";"

	var response string
	collector := colly.NewCollector()
	collector.OnResponse(func(r *colly.Response) {
		response = string(r.Body)
	})
	collector.Visit(chapterURL.String())

	if len(response) == 0 {
		log.Fatal("Could not find pages urls at specified MangaPark URL")
	}

	response = strings.SplitAfter(response, splitAfterSep)[1]
	response = strings.Split(response, splitBeforeSep)[0]

	return getImageUrlsFromJSON([]byte(response))
}

func getImageUrlsFromJSON(jsonResponse []byte) []string {
	var pagesJSON []map[string]interface{}
	err := json.Unmarshal(jsonResponse, &pagesJSON)

	if err != nil {
		log.Fatal("Could not parse pages JSON at specified MangaPark URL")
	}

	var imageUrls []string
	for _, pageObj := range pagesJSON {
		imageUrls = append(imageUrls, pageObj["u"].(string))
	}
	return imageUrls
}
