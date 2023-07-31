package main

import (
	"fmt"

	"github.com/brewingweasel/lithuanian-info-parser-lkz/parsing"
	"github.com/brewingweasel/lithuanian-info-parser-lkz/scraping"
	"github.com/charmbracelet/log"
)

func main() {
	log.SetLevel(log.DebugLevel)
	id := scraping.GetIdOfWord("mušti")
	log.Info(id)
	lkzContents := scraping.GetWordDetails(id)
	fmt.Println(parsing.GetVerbInfo(lkzContents))
}
