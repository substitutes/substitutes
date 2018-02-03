package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"github.com/fronbasal/vertretungsplan/structs"
)

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

func LoadCredentials() structs.Credentials {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatal("Failed to read config file!")
	}
	var c structs.Credentials
	json.Unmarshal(b, &c)
	return c
}

func IServLogin(username, password string) (error, bool) {
	body := strings.NewReader(`_username=` + username + `&_password=` + password)
	req, err := http.NewRequest("POST", "https://steinbart-gym.eu/iserv/login_check", body)
	if err != nil {
		return err, false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err, false
	}
	defer resp.Body.Close()
	location, err := resp.Location()
	if err != nil {
		return err, false
	}
	return nil, location.Path == "/iserv"
}
