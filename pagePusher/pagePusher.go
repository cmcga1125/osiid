package pagePusher

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/storage"
)

func PushHtml(path string, html []byte, bucket string) {
	object := strings.TrimPrefix(path, "/") + "index.html"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("storage.NewClient: %v", err)
	}
	defer client.Close()

	buf := bytes.NewBuffer(html)

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err = io.Copy(wc, buf); err != nil {
		log.Fatal("io.Copy: %v", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		log.Fatal("Writer.Close: %v", err)
	}
	fmt.Printf("%v uploaded to %v.\n", object, bucket)
}