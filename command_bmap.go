package main

import "pokedexcli/internal/pokeapi"

func commandBMap(client *pokeapi.Client) error{
	result, err := client.FetchLocation()
	if err != nil{
		return err
	}
	pokeapi.PrintMap(result)
	client.MovePreviousPageMap(*result)
	return nil
}