package PageGetter

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	// "path"
)

func GetPage(httpUrl string) {
	u, err := url.Parse(httpUrl)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get(httpUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	path := "./tmp" + u.Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}
	error := os.WriteFile(path + "index.html", body, 0777)
	if error != nil {
		log.Fatal(error)
	}
}