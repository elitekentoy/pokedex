package repl

func commandPokedex(config *Config, _ string) error {
	if len(config.Pokedex) == 0 {
		println("You don't have any pokemons in your pokedex, You gotta catch 'em all!")
		return nil
	}
	println("Your pokedex: ")
	for _, pokemon := range config.Pokedex {
		println("- ", pokemon.Name)
	}
	return nil
}
