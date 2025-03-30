package repl

import (
	"github.com/elitekentoy/pokedexcli/internal/pokeapi"
)

func commandExplore(config *Config, area string) error {
	url := pokeapi.Pokeapi_location_area_url + "/" + area

	locationResp, err := config.PokeApiClient.LocationDetails(&url)
	if err != nil {
		return err
	}

	println("Exploring " + area + "...")
	println("Found Pokemon: ")
	for _, encounter := range locationResp.PokemonEncounters {
		pokemon := encounter.Pokemon
		println(pokemon.Name)
	}

	return nil
}
