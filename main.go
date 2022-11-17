package main

import (
	"osiid.io/pageGetter"
)

func main() {
	pages := []string {
		"https://www.cnet.com/roadshow/news/2023-toyota-prius-teaser-livestream-debut/",
		"https://www.cnet.com/",
		}
	for _, e := range pages {
		PageGetter.GetPage(e)
	}
	gcs()
}