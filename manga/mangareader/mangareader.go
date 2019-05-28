package mangareader

import (
	"net/url"

	"github.com/gocolly/colly"
)

const pageNumberSelector string = "#pageMenu > option"
const imageSelector string = "#img"

var collector = colly.NewCollector()

func FetchPageImages(chapterURL *url.URL) []string {
	pageNumbers := fetchPageNumbers(chapterURL)
	imagesUrls := fetchImageUrls(chapterURL, pageNumbers)
	return fetchManagerReaderImages(imagesUrls)
}

func fetchPageNumbers(chapterURL *url.URL) []string {
	var pageNumbers []string
	collector.OnHTML(pageNumberSelector, func(element *colly.HTMLElement) {
		pageNumbers = append(pageNumbers, element.Text)
	})
	collector.Visit(chapterURL.String())

	collector.OnHTMLDetach(pageNumberSelector)
	return pageNumbers
}

func fetchImageUrls(chapterURL *url.URL, pageNumbers []string) []string {
	imagesUrls := make([]string, 0, len(pageNumbers))
	collector.OnHTML(imageSelector, func(element *colly.HTMLElement) {
		imagesUrls = append(imagesUrls, element.Attr("src"))
	})
	for _, e := range pageNumbers {
		url := chapterURL.String() + "/" + e
		collector.Visit(url)
	}

	collector.OnHTMLDetach(imageSelector)
	return imagesUrls
}

func fetchManagerReaderImages(imagesUrls []string) []string {
	images := make([]string, 0, len(imagesUrls))

	collector.OnResponse(func(response *colly.Response) {
		images = append(images, string(response.Body))
	})

	for _, e := range imagesUrls {
		collector.Visit(e)
	}

	return images
}
