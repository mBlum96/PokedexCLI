package main

import (
	"pokedexcli/internal/pokeapi"
)

func commandMap(client *pokeapi.Client) error{
	result, err := client.FetchLocation()
	if err != nil{
		return err
	}
	pokeapi.PrintMap(result)
	client.MoveNextPageMap(*result)
	return nil
}