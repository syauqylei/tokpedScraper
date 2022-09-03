package main

import (
	"fmt"

	"github.com/syauqylei/tokpedScraper/docLoader"
	"github.com/syauqylei/tokpedScraper/docParser"
)

func main() {
	c := docLoader.DocLoaderCtx{}
	htmlCtx := docLoader.GetDocs(&c)
	phones := docParser.ParsePhones(htmlCtx)
	docParser.SaveToCsv(phones, "tokped100.csv")
	fmt.Println(len(phones))
}
