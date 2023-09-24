package main

var myConfig = &config{nextURL: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20", previousURL: "Your journey has only just begun!"}

type config struct {
	nextURL     string
	previousURL string
}
