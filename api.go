package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getNextMapPage(offset int) {
	link := fmt.Sprintf("http://pokeapi.co/api/v2/location/?offset=%v", offset)
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

	fmt.Printf("%s", body)
}