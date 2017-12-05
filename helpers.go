package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func request(url string) (*http.Response, error) {
	creds := loadCredentials()
	req, err := http.NewRequest("GET", creds.Host+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.Username, creds.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func loadCredentials() Credentials {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatal("Failed to read config file!")
	}
	var c Credentials
	json.Unmarshal(b, &c)
	return c
}
