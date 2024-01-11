package main

import (
	"pokedexcli/internal/pokeapi"
	"pokedexcli/pokecache"
	"time"
)

type cliCommand struct{
    name string
    description string
    callback func() error
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
            callback: func() error {
                return commandMap(client)
            },
        },
        "bmap":{
            name: "bmap",
            description: "Show the previous map",
            callback: func() error {
                return commandBMap(client)
            },
        },
    }
}

func main(){
    const interval = time.Minute
    myCache:= pokecache.NewCache(interval)
    repl()
}

