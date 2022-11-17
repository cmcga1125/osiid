package main

import (
	"log"
	"net/url"

	"osiid.io/pageGetter"
	"osiid.io/pagePusher"
)

func main() {
	bucket := "test-bucket"
	pages := []string {
		"https://www.your-page.com/about/",
		"https://www.your-page.com/",
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