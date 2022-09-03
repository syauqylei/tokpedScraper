package docLoader

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DocLoaderCtx struct {
	Url   string
	Docs  [25]*goquery.Document
	Order int
}

const (
	tokpedPhoneURL string = "https://www.tokopedia.com/p/handphone-tablet/handphone"
	userAgent      string = "Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0"
	order          int    = 23
)

func initDefaultConfig(c *DocLoaderCtx) {
	c.Url = tokpedPhoneURL
	c.Order = order
}

func ReqDoc(url string) *goquery.Document {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	html, _ := ioutil.ReadAll(resp.Body)
	htmlString := string(html)
	reader := strings.NewReader(htmlString)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func GetUrl(c *DocLoaderCtx, page int) string {
	url := c.Url + "?ob=" + strconv.Itoa(c.Order) + "?page=" + strconv.Itoa(page+1)
	return url
}

func GetDocs(c *DocLoaderCtx) *DocLoaderCtx {
	initDefaultConfig(c)

	for i := 0; i < 25; i++ {
		curUrl := GetUrl(c, i)
		exDoc := ReqDoc(curUrl)
		c.Docs[i] = exDoc
	}
	return c
}
