package scraping

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gocolly/colly"
)

func getUrlForWordId(word string) string {
	return "http://lkz.lt/Zodziai.asp?txtZodis=" + word + "&nrLn=-1&nrLe=-1&zdId=-1&tstMode=0"
}

func getUrlFromId(id string) string {
	return "http://lkz.lt/Zd" + id[0:2] + "/" + id + ".htm?"
}

func GetIdOfWord(word string) string {
	c := colly.NewCollector()
	var wordId string

	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting: %s", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatal(err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Debugf("Page visited: %s", r.Request.URL)
	})

	c.OnHTML("#Sarasas div", func(e *colly.HTMLElement) {
		log.Infof("Found %s", e.Text)
		text := strings.Trim(e.Text, "1234567890")
		if text == word && len(wordId) == 0 {
			wordId = strings.TrimSuffix(strings.TrimPrefix(e.Attr("id"), "d"), "000")
		}
	})

	c.Visit(getUrlForWordId(word))
	return wordId
}

func GetWordDetails(id string) string {
	c := colly.NewCollector()

	var wordInfo string
	log.SetLevel(log.InfoLevel)

	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting: %s", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatal(err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Debugf("Page visited: %s", r.Request.URL)
	})

	c.OnHTML(".az", func(e *colly.HTMLElement) {
		log.Debug(e)
		wordInfo = e.Text
	})

	c.OnScraped(func(r *colly.Response) {
		log.Debugf("Page scraped: %s", r.Request.URL)
	})

	c.Visit(getUrlFromId(id))
	return wordInfo
}
