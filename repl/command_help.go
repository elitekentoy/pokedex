package repl

import (
	"fmt"
)

func commandHelp(config *Config, _ string) error {
	intro := "Welcome to the Pokedex!\nUsage:\n\n"
	usages := ""

	for key, command := range GetCommands() {
		usages += key + ": " + command.description + "\n"
	}

	full := intro + usages

	fmt.Println(full)
	return nil
}
