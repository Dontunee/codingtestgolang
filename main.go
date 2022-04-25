package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

func main() {
	StartServer()
}

func StartServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/listings", getListingsHandler)
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	data := []byte("Hello, World!")
	w.Write(data)
}

func getListingsHandler(w http.ResponseWriter, req *http.Request) {
	listings := GetListings()

	json, err := json2.Marshal(listings)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(200)
	w.Write(json)
}

func GetListings() []Listing {
	var listings []Listing
	listings = []Listing{
		{
			ID:            1,
			StreetAddress: "toronto",
			Price:         1000000,
			Status:        "for sale",
		},
		{
			ID:            2,
			StreetAddress: "oshawa",
			Price:         800000,
			Status:        "for sale",
		},
	}
	return listings
}
