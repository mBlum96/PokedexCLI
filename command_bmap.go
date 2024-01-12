package main

import "pokedexcli/internal/pokeapi"

func commandBMap(client *pokeapi.Client) error{
	result, err := client.FetchLocation()
	if err != nil{
		return err
	}
	client.MovePreviousPageMap(*result)
	pokeapi.PrintMap(result)
	return nil
}