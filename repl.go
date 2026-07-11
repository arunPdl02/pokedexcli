package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/arunPdl02/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPIClient    pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := reader.Scan()
		if !ok {
			break
		}
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		} else {
			args = nil
		}

		command_registry := getCommandRegistry()
		command, ok := command_registry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg, args...); err != nil {
			fmt.Printf("Error with command %s: %v", command.name, err)
		}
	}
}

func cleanInput(text string) []string {
	result := []string{}
	text = strings.ToLower(text)

	start := -1

	for i, char := range text {
		if char != ' ' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				result = append(result, text[start:i])
				start = -1 // Reset state
			}
		}
	}

	if start != -1 {
		result = append(result, text[start:])
	}

	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show the list of pokemon in the requested location. Pass the location as an argument",
			callback:    commandExplore,
		},
	}
}
