package main

import (
	"math/rand"
	"pokedexcli/internal/pokeapi"
	"time"
)

func commandCatch(c *pokeapi.Client, pokemonName string) error{
	pokemonEndpoint := pokeapi.DefaultBaseUrl + "pokemon"
	pokemonData,err := c.FetchPokemonInformation(pokemonEndpoint+"/"+pokemonName)
	if(err!=nil){
		return err
	}
	tryCatchingPokemon(*pokemonData, c)
	return nil
}


func tryCatchingPokemon(pokemonData pokeapi.PokemonInfo, c *pokeapi.Client ){
	//because pidgey is considered an easy pokemon to catch, we will use it as a base
	//for the catch rate, it's base experience is 50, so we will use that - halfed as a base for the catch rate
	//we will use the formula: catchRate = (3*maxHP - 2*currentHP)*catchRate*ballBonus/(3*maxHP) when we expend the program

	//for now we ask for a random number from 1 to the base experience of the pokemon
	//if the number is lower than the base pidgey exp, the pokemon is caught

	pokemonBaseExp := pokemonData.BaseExperience
	randomNumber := rand.Intn(pokemonBaseExp) + 1
	print("Throwing pokeball at " + pokemonData.Name + ".")
	time.Sleep(500 * time.Millisecond)
	print(".")
	time.Sleep(500 * time.Millisecond)
	print(".\n")
	if(randomNumber<pokeapi.PidgeyBaseExperience){
		c.AddToPokeDex(pokemonData)
		println(pokemonData.Name + " was caught!")
	}else{
		println(pokemonData.Name + " escaped!")
	}
}