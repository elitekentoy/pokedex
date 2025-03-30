package repl

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, name string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if config.CatchTries[name] >= 5 {
		println("Catch Tries Exceeded...")
		return nil
	}

	config.CatchTries[name]++

	_, exist := config.Pokedex[name]
	if exist {
		println(name, " already exists in the Pokedex")
		return nil
	}

	pokemon, err := config.PokeApiClient.GetPokemonDetails(name)
	if err != nil {
		return err
	}

	if calculateChances(pokemon.BaseExperience) >= 60 {
		config.Pokedex[name] = pokemon
		println(name + " was caught!")
	} else {
		println(name + " escaped!")
	}

	return nil
}

func calculateChances(baseXp int) int {
	chances := rand.Intn(100)

	if baseXp > 100 {
		chances -= 20
	} else {
		chances += 20
	}

	return chances
}
