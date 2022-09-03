package helper

import (
	"log"
	"regexp"
)

func TrimSymbols(s string) string {
	re, err := regexp.Compile(`(Rp|\.)`)
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllString(s, "")
	return s
}
