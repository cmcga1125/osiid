package main

import (
	"log"
	"net/url"

	"osiid.io/pageGetter"
	"osiid.io/pagePusher"
)

func main() {
	bucket := "corey-test-bucket"
	pages := []string {
		"https://www.cnet.com/roadshow/news/2023-toyota-prius-teaser-livestream-debut/",
		"https://www.cnet.com/",
		}
	for _, e := range pages {
		uri, err := url.Parse(e)
		if err != nil {
			log.Fatal(err)
		}
		html := PageGetter.GetHtml(e)
		pagePusher.PushHtml(uri.Path, html, bucket)
	}
}