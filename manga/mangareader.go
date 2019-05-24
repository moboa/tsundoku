package manga

import (
	"container/list"
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
)

// PrintChapterPages prints all page numbers of the chapter at the given URL.
func PrintChapterPages(chapterURL *url.URL) {
	pageNumbers := list.New()
	collector := colly.NewCollector()
	collector.OnHTML("#pageMenu > option", func(element *colly.HTMLElement) {
		pageNumbers.PushBack(element.Text)
	})
	collector.Visit(chapterURL.String())

	collector.OnHTMLDetach("#pageMenu > option")
	imageUrls := list.New()
	collector.OnHTML("#img", func(element *colly.HTMLElement) {
		imageUrls.PushBack(element.Attr("src"))
	})
	for e := pageNumbers.Front(); e != nil; e = e.Next() {
		url := chapterURL.String() + "/" + e.Value.(string)
		collector.Visit(url)
	}

	collector.OnHTMLDetach("#img")
	images := list.New()
	collector.OnResponse(func(response *colly.Response) {
		images.PushBack(response.Body)
		fmt.Println(response.Body)
	})
	for e := imageUrls.Front(); e != nil; e = e.Next() {
		collector.Visit(e.Value.(string))
	}
}
