package main

import (
	"pokedexcli/internal/pokeapi"
)

func commandBMap(client *pokeapi.Client) error{
	result, err := client.FetchLocation(pokeapi.PREV)
	if err != nil{
		return err
	}
	pokeapi.PrintMap(result)
	return nil
}