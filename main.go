package main

import (
	"errors"
	"pokedexcli/internal/pokeapi"
)

type cliCommand struct{
    name string
    description string
    callback func([]string) error
}

var commands map[string]cliCommand

func init() {
    client:= pokeapi.NewClient()
    commands = map[string]cliCommand{
        "help": {
            name: "help",
            description: "Show help",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the program",
            callback: commandExit,
        },
        "map":{
            name: "map",
            description: "Show the map",
            callback: func(params []string) error {
                return commandMap(client)
            },
        },
        "bmap":{
            name: "bmap",
            description: "Show the previous map",
            callback: func(params []string) error {
                return commandBMap(client)
            },
        },
        "explore":{
            name: "explore",
            description: "Explore the chosen location",
            callback: func(params []string) error{
                if len(params) < 1{
                    return errors.New("No location provided for 'explore' command")
                }
                location:=params[0]
                return commandExplore(client, location)
            },
        },
    }
}

func main(){
    repl()
}

