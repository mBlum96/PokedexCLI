package main

import (
	"pokedexcli/internal/pokeapi"
)

func commandMap(client *pokeapi.Client) error{
	result, err := client.FetchLocation(pokeapi.NEXT)
	if err != nil{
		return err
	}
	pokeapi.PrintMap(result)
	return nil
}