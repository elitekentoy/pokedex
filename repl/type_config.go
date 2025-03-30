package repl

import "github.com/elitekentoy/pokedexcli/internal/pokeapi"

type Config struct {
	Next          *string
	Previous      *string
	PokeApiClient pokeapi.Client
	Pokedex       map[string]pokeapi.Pokemon
	CatchTries    map[string]int
}
