package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mapResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	resp, err := http.Get(cfg.nextURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	body := mapResponse{}

	if err := json.Unmarshal(respBody, &body); err != nil {
		return err
	}

	for _, result := range body.Results {
		fmt.Println(result.Name)
	}

	// assign config values for next map command
	cfg.nextURL = body.Next
	if body.Previous != nil {
		cfg.previousURL = *body.Previous
	} else {
		cfg.previousURL = ""
	}

	return nil
}
