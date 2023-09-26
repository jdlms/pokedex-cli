package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandBack(cfg *config) error {

	if cfg.previousURL != "Your journey has only just begun!" {

		resp, err := http.Get(cfg.previousURL)
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

		cfg.nextURL = body.Next
		if body.Previous != nil {
			cfg.previousURL = *body.Previous
		}

	} else {
		fmt.Println(cfg.previousURL)
	}

	return nil
}
