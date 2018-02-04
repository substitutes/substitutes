package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fronbasal/substitutes/structs"
)

// Request helper function for making a web request
func Request(url string) (*http.Response, error) {
	credentials := LoadCredentials()
	req, err := http.NewRequest("GET", credentials.Host+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(credentials.Username, credentials.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LoadCredentials helper functions for loading credentials.json
func LoadCredentials() structs.Credentials {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatal("Failed to read config file!")
	}
	var c structs.Credentials
	json.Unmarshal(b, &c)
	return c
}

// IServLogin for authenticating against IServ
func IServLogin(username, password string) (bool, error) {
	body := strings.NewReader(`_username=` + username + `&_password=` + password)
	req, err := http.NewRequest("POST", "https://steinbart-gym.eu/iserv/login_check", body)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	location, err := resp.Location()
	if err != nil {
		return false, err
	}
	return location.Path == "/iserv", nil
}
