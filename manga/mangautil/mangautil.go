package mangautil

import (
	"github.com/gocolly/colly"
)

// FetchImages downloads images using provided URLs
func FetchImages(collector *colly.Collector, imagesUrls []string) []string {
	images := make([]string, 0, len(imagesUrls))

	collector.OnResponse(func(response *colly.Response) {
		images = append(images, string(response.Body))
	})

	for _, e := range imagesUrls {
		collector.Visit(e)
	}

	return images
}
