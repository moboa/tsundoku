package scrape

import (
	"net/url"

	"github.com/gocolly/colly"
)

const pageNumberSelector string = "#pageMenu > option"
const imageSelector string = "#img"

// FetchPageImages returns a list containing images of the chapter at specified URL
func FetchPageImages(chapterURL *url.URL) []string {
	var pageNumbers []string
	collector := colly.NewCollector()
	collector.OnHTML(pageNumberSelector, func(element *colly.HTMLElement) {
		pageNumbers = append(pageNumbers, element.Text)
	})
	collector.Visit(chapterURL.String())

	collector.OnHTMLDetach(pageNumberSelector)
	imageUrls := make([]string, 0, len(pageNumbers))
	collector.OnHTML(imageSelector, func(element *colly.HTMLElement) {
		imageUrls = append(imageUrls, element.Attr("src"))
	})
	for _, e := range pageNumbers {
		url := chapterURL.String() + "/" + e
		collector.Visit(url)
	}

	collector.OnHTMLDetach(imageSelector)
	images := make([]string, 0, len(imageUrls))
	collector.OnResponse(func(response *colly.Response) {
		images = append(images, string(response.Body))
	})
	for _, e := range imageUrls {
		collector.Visit(e)
	}

	return images
}
