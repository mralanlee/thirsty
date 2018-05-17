package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func mapURL(env string) string {
	url := fmt.Sprintf("http://url%s/", env)
	return url
}

// CreateSession generates the session needed
func createSession(env string) []byte {
	resp, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}

// TODO refactor
func validateSession(env string) []byte {
	resp, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}
