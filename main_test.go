package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	go StartServer()

	t.Run("Get Listings Test", func(t *testing.T) {
		response, err := http.Get("http://localhost:8080/listings")
		if err != nil {
			t.Errorf("Test failed with error response from endpoint")
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject []Listing
		json.Unmarshal(responseData, &responseObject)

		if len(responseObject) <= 0 {
			t.Errorf("Test failed with no data in response body")
		}

		if responseObject[0].ID != 1 && responseObject[0].StreetAddress != "toronto" {
			t.Errorf("Test failed due to wrong data in response body")
		}

	})

}
