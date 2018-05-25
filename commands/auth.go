package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type authURL struct {
	create, validate, session string
}

type createResponse struct {
	sessionKey string
}

// CreateResponse for capturing createResponse type
type CreateResponse struct {
	createResponse
}

func mapURL(env string) map[string]string {
	urls := make(map[string]string)
	urls["create"] = fmt.Sprintf("https://%s", env)
	urls["validate"] = fmt.Sprintf("https://%s", env)
	urls["cookie"] = fmt.Sprintf("https://%s", env)
	return urls
}

// CreateSession generates the session needed
func createSession(env string) string {
	resp, err := http.Get("http://www.example.com")

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var result createResponse
	json.Unmarshal(body, &result)

	defer resp.Body.Close()

	return result.sessionKey
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

func getCookie(env string) []byte {
	resp, err := http.Get("http://www.example.com/id2")

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}

// UnmarshalJSON parses json
func (d *CreateResponse) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.createResponse)
}

// GetKEY will fetch the field
func (d *CreateResponse) GetKEY(field string) string {
	return d.createResponse.sessionKey
}
