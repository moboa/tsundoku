package mangareader

import (
	"github.com/gocolly/colly"
	"github.com/moboa/tsundoku/manga/mangautil"
)

const pageNumSelector string = "#pageMenu > option"
const imageSelector string = "#img"

// FetchPageImages returns a list containing images of the chapter at specified MangaReader URL
func FetchPageImages(collector *colly.Collector, chapterURL string) []string {
	pageNumbers := fetchPageNumbers(collector, chapterURL)
	imagesUrls := fetchImageUrls(collector, chapterURL, pageNumbers)
	return mangautil.FetchImages(collector, imagesUrls)
}

func fetchPageNumbers(collector *colly.Collector, chapterURL string) []string {
	var pageNumbers []string
	collector.OnHTML(pageNumSelector, func(element *colly.HTMLElement) {
		pageNumbers = append(pageNumbers, element.Text)
	})
	collector.Visit(chapterURL)

	collector.OnHTMLDetach(pageNumSelector)
	return pageNumbers
}

func fetchImageUrls(collector *colly.Collector, chapterURL string, pageNumbers []string) []string {
	imagesUrls := make([]string, 0, len(pageNumbers))
	collector.OnHTML(imageSelector, func(element *colly.HTMLElement) {
		imagesUrls = append(imagesUrls, element.Attr("src"))
	})
	for _, e := range pageNumbers {
		url := chapterURL + "/" + e
		collector.Visit(url)
	}

	collector.OnHTMLDetach(imageSelector)
	return imagesUrls
}
