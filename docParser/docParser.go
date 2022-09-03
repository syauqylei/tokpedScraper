package docParser

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/syauqylei/tokpedScraper/docLoader"
	"github.com/syauqylei/tokpedScraper/helper"
	"github.com/syauqylei/tokpedScraper/phone"
)

const top100 = 100

func getPrice(s *goquery.Selection) int64 {
	price := s.Find(".css-o5uqvq").Text()
	price = helper.TrimSymbols(price)

	n, err := strconv.ParseInt(price, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func getName(s *goquery.Selection) string {
	return s.Find(".css-1bjwylw").Text()
}

func getStore(s *goquery.Selection) string {
	return s.Find(".css-1kr22w3").Eq(1).Text()
}

func getRating(s *goquery.Selection) int {
	// descUrl, _ := s.Find(".css-89jnbj").Attr("href")
	stars := s.Find(".css-153qjw7").
		Children().
		First().Children().Filter("img").Nodes
	// ratingParsed, _ := strconv.Atoi(rating)
	return len(stars)
}

func getImageLink(s *goquery.Selection) string {
	src, _ := s.Find(".fade").Attr("src")
	return src
}

func getDescription(s *goquery.Selection) string {
	descUrl, _ := s.Find(".css-89jnbj").Attr("href")
	doc := docLoader.ReqDoc(descUrl)
	desc := doc.Find(".eytdjj01").Text()
	return desc
}

func ParsePhones(c *docLoader.DocLoaderCtx) []phone.Phone {
	counter := 0
	phones := make([]phone.Phone, 0)
	for i := 0; i < 25; i++ {
		c.Docs[i].
			Find(".css-bk6tzz").
			EachWithBreak(func(i int, s *goquery.Selection) bool {
				counter++
				if counter == 100 {
					return false
				}

				curObj := new(phone.Phone)
				curObj.Name = getName(s)
				curObj.Store = getStore(s)
				curObj.Price = getPrice(s)
				curObj.ImageLink = getImageLink(s)
				curObj.Description = getDescription(s)
				curObj.Rating = getRating(s)
				if curObj.Description == "" {
					counter--
				} else {
					phones = append(phones, *curObj)
				}
				return true
			})
	}

	return phones
}
