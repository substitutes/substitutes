package main

import (
	"testing"
	"io/ioutil"
	"github.com/substitutes/substitutes/structs"
	"encoding/json"
	"net/http"
)

func TestClassOutputLength(t *testing.T) {
	// Request API
	response, err := http.Get("http://localhost:5000/api/c/11")
	if err != nil {
		t.Error(err)
	}

	defer response.Body.Close()

	// Marshal into body

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	var v struct{ Data []structs.Substitute `json:"substitutes"` } // Only marshal substitutes
	if err := json.Unmarshal(bytes, &v); err != nil {
		t.Logf("Data: %v", string(bytes))
		t.Error(err)
	}

	if len(v.Data) == 0 {
		// Fail if empty
		t.Log("Length of classes is 0 - something is wrong!")
		t.Fail()
	}
}
