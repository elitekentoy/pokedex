package repl

func commandInspect(config *Config, name string) error {
	pokemon, exists := config.Pokedex[name]
	if !exists {
		println("Cannot inspect ", name, " because it does not exist in the pokedex")
		return nil
	}

	println("Name: ", pokemon.Name)
	println("Height: ", pokemon.Height)
	println("Weight: ", pokemon.Weight)
	println("Stats: ")
	for _, stat := range pokemon.Stats {
		println("-", stat.Stat.Name, ": ", stat.BaseStat)
	}

	println("Types: ")
	for _, pokeType := range pokemon.Types {
		println("- ", pokeType.Type.Name)
	}

	return nil
}
