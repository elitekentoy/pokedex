package repl

import (
	"errors"
	"fmt"
)

func commandMapf(config *Config, _ string) error {
	locationResp, err := config.PokeApiClient.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationResp.Next
	config.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(config *Config, _ string) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.PokeApiClient.ListLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationResp.Next
	config.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
