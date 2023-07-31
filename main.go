package main

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/log"
	"github.com/gocolly/colly"
	"golang.org/x/exp/slices"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var LITH_SPECIAL_CHARS = []rune{'ą', 'č', 'ę', 'ė', 'į', 'š', 'ų', 'ū', 'ž'}

func getUrlForWordId(word string) string {
	return "http://lkz.lt/Zodziai.asp?txtZodis=" + word + "&nrLn=-1&nrLe=-1&zdId=-1&tstMode=0"
}

func getUrlFromId(id string) string {
	return "http://lkz.lt/Zd" + id[0:2] + "/" + id + ".htm?"
}

func getIdOfWord(word string) string {
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

func getWordDetails(id string) string {
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

func getVerbInfo(grammaticalInfo string) [3]string {
	verbData := strings.Split(removeAccentuation(grammaticalInfo), ",")
	log.Debug(verbData)
	return [3]string{verbData[0], verbData[1], strings.SplitN(verbData[2], " ", 3)[1]}

}

func main() {
	log.SetLevel(log.DebugLevel)
	id := getIdOfWord("mušti")
	log.Info(id)
	details := getWordDetails(id)
	log.Info(createVerbVals(getVerbInfo(details)))
}

func getAccents(r rune) bool {
	if slices.Contains(LITH_SPECIAL_CHARS, r) {
		return false
	}
	return unicode.Is(unicode.Mn, r)
}

// Removes actual accentuation marks, not standard Lithuanian characters
// ie: pralaũkęs: pralaukęs
func removeAccentuation(word string) string {
	t := transform.Chain(runes.Remove(runes.Predicate(getAccents)), norm.NFC)
	result, _, _ := transform.String(t, word)
	return result
}

func createVerbVals(verbInfo [3]string) [3]string {
	log.Info(verbInfo)
	verbInfo[1] = strings.TrimSpace(verbInfo[1])
	verbInfo[2] = strings.TrimSpace(verbInfo[2])
	if strings.HasPrefix(verbInfo[1], "-") {
		log.Debug("Using prefixed")
		if len(verbInfo[1]) == 2 {
			removedVerbEnding := strings.TrimSuffix(verbInfo[0], "ti")
			if strings.HasSuffix(removedVerbEnding, "y") || strings.HasSuffix(removedVerbEnding, "ė") {
				removedVerbEnding = removedVerbEnding[:len(removedVerbEnding)-2]
			}
			verbInfo[1] = removedVerbEnding + verbInfo[1][1:]
			verbInfo[2] = removedVerbEnding + verbInfo[2][1:]
		}
		for i := len(verbInfo[0]) - 2; i > 0; i-- {
			if verbInfo[0][i:i+2] == verbInfo[1][1:3] {
				verbInfo[1] = verbInfo[0][:i] + verbInfo[1][1:]
				verbInfo[2] = verbInfo[0][:i] + verbInfo[2][1:]
				break
			}
		}
	}
	return verbInfo
}
