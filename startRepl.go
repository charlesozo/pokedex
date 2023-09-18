package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/charlesozo/pokedex/internal/api"

)

type config struct{
	pokeapiClient api.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
}

func startRepl(cfg *config){
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

	
		commandName := words[0]
		var inputLocation string
		if len(words)>1{
			inputLocation = words[1]
			
		}


		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, inputLocation)
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
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
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
			description: "List of some Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List of previous Locations",
			callback:    commandMapb,
		},
		"explore":{
			name:        "explore <area_name>",
			description: "List of pokemons in the area",
			callback:    commandExplore,
		},
	}
}