package mangastream

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/moboa/tsundoku/manga/mangautil"
)

const pageNumSelector string = "div.btn-reader-page > ul.dropdown-menu > li"
const imageSelector string = "#manga-page"
const splitAfterSep string = "Last Page ("
const splitBeforeSep string = ")"

// FetchPageImages returns a list containing images of the chapter at specified MangaReader URL
func FetchPageImages(collector *colly.Collector, chapterURL string) []string {
	lastPageNum := fetchLastPageNumber(collector, chapterURL)
	imagesUrls := fetchImageUrls(collector, chapterURL, lastPageNum)
	return mangautil.FetchImages(collector, imagesUrls)
}

func fetchLastPageNumber(collector *colly.Collector, chapterURL string) int {
	lastPageNum := -1
	collector.OnHTML(pageNumSelector, func(element *colly.HTMLElement) {
		if !strings.Contains(element.Text, "Last") {
			return
		}

		pageNum := strings.SplitAfter(element.Text, splitAfterSep)[1]
		pageNum = strings.Split(pageNum, splitBeforeSep)[0]
		lastPageNum, _ = strconv.Atoi(pageNum)
	})
	collector.Visit(chapterURL)

	collector.OnHTMLDetach(pageNumSelector)

	if lastPageNum == -1 {
		log.Fatal("Could not find the number of the last page of the chapter at the specified Mangastream URL")
	}
	return lastPageNum
}

func fetchImageUrls(collector *colly.Collector, chapterURL string, lastPageNum int) []string {
	imagesUrls := make([]string, 0, lastPageNum)
	collector.OnHTML(imageSelector, func(element *colly.HTMLElement) {
		imagesUrls = append(imagesUrls, element.Attr("src"))
	})
	for i := 1; i <= lastPageNum; i++ {
		url := chapterURL + "/" + strconv.Itoa(i)
		collector.Visit(url)
	}

	collector.OnHTMLDetach(imageSelector)
	return imagesUrls
}
