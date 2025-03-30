package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	url := Pokeapi_pokemon_url + "/" + pokemonName

	if val, ok := client.cache.Get(url); ok {
		cachedData := Pokemon{}
		err := json.Unmarshal(val, &cachedData)
		if err != nil {
			return Pokemon{}, err
		}
		return cachedData, nil
	}

	// Create Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// Execute Request
	res, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonDetails := Pokemon{}
	err = json.Unmarshal(data, &pokemonDetails)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(url, data)

	return pokemonDetails, err
}
