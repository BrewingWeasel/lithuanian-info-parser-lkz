package main

import (
	"fmt"

	"github.com/brewingweasel/lithuanian-info-parser-lkz/parsing"
	"github.com/brewingweasel/lithuanian-info-parser-lkz/scraping"
	"github.com/charmbracelet/log"
)

func main() {
	word := "kąsnis"
	log.SetLevel(log.WarnLevel)
	id := scraping.GetIdOfWord(word)
	log.Info(id)
	lkzContents := scraping.GetWordDetails(id)
	fmt.Println(parsing.GetGenderDecl(lkzContents, word))
}
