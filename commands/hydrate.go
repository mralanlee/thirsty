package commands

import (
	"io/ioutil"
	"net/http"
)

var baseURL = "/_changes?filter=sync_gateway/bychannel&channels="

func getChangesFeed(url string, ch chan<- string) []byte {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
