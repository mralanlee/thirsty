package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var bucket = "asset360"

func getChangesFeed(url string, ch chan<- string) []byte {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)

	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)

	return body
}
