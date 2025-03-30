package main

import (
	"time"

	"github.com/elitekentoy/pokedexcli/internal/pokeapi"
	"github.com/elitekentoy/pokedexcli/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &repl.Config{
		PokeApiClient: pokeClient,
		Pokedex:       map[string]pokeapi.Pokemon{},
		CatchTries:    map[string]int{},
	}

	repl.StartRepl(config)
}
