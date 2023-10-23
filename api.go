package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func getNextMapPage(link string) locations {
	res, err := http.Get(link)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	structuredResponse := locations{}

	err = json.Unmarshal([]byte(body), &structuredResponse)

	if err != nil {
		log.Fatal(err)
	} 

	return structuredResponse
}