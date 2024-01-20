package main

import (
	"errors"
	"pokedexcli/internal/pokeapi"
)

type cliCommand struct{
    name string
    description string
    callback func(*pokeapi.Client ,[]string) error
}

func callbackWithParams(command func(*pokeapi.Client, string) error) func(*pokeapi.Client, []string) error{
    return func(client *pokeapi.Client, params []string) error{
        if(len(params)>1){
            return errors.New("Too many parameters")
        }else if(len(params)<1){
            return errors.New("Too few parameters")
        }
        return command(client, params[0])
    }
}

func callbackNoParams(command func(*pokeapi.Client) error) func(*pokeapi.Client, []string) error{
    return func(client *pokeapi.Client, params []string) error{
        return command(client)
    }
}

var commands map[string]cliCommand

func main(){
    client:= pokeapi.NewClient()
    commands = map[string]cliCommand{
        "help": {
            name: "help",
            description: "Show help",
            callback: callbackNoParams(commandHelp),
        },
        "exit": {
            name: "exit",
            description: "Exit the program",
            callback: callbackNoParams(commandExit),
        },
        "map":{
            name: "map",
            description: "Show the map",
            callback: callbackNoParams(commandMap),
        },
        "bmap":{
            name: "bmap",
            description: "Show the previous map",
            callback: callbackNoParams(commandBMap),
        },
        "explore":{
            name: "explore",
            description: "Explore the chosen location",
            callback: callbackWithParams(commandExplore),
        },
        "catch":{
            name: "catch",
            description: "Attempt to catch the chosen pokemon",
            callback: callbackWithParams(commandCatch),
        },
        "inspect":{
            name: "inspect",
            description: "Get information regarding the chosen pokemon",
            callback: callbackWithParams(commandInspect),
        },
        "pokedex":{
            name: "pokedex",
            description: "List caught pokemon",
            callback: callbackNoParams(commandPokedex),
        },
    }
    repl(client)
}

