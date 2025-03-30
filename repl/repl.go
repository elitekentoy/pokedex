package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		argument := ""

		if len(words) == 2 {
			argument = words[1]
		}

		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(config, argument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Displays the previous 20 location ares in Pokemon world (if applicable)",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the pokemons of the specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throws a pokeball to try catching the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects an existing pokemon in the pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Checks all the pokemons in the pokedex",
			callback:    commandPokedex,
		},
	}
}
