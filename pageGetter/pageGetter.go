package PageGetter

import (
	"io"
	"log"
	"net/http"
)

func GetPage(httpUrl string) []byte {
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
	return body
}