package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
)

func commandExit(c *pokeapi.Client) error{
    fmt.Println("Exiting the pokedex...")
    os.Exit(0)
    return nil
}