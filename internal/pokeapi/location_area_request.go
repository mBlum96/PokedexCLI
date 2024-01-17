package pokeapi

import (
	"fmt"
)


func PrintMap(response *LocationResponse){
    for _,area := range(response.Locations){
        fmt.Printf("%s\n", area.Name)
    }
}

func PrintPokemon(response *pokemonEncountered){
    for _,pokemon := range(response.PokemonEncounters){
        fmt.Printf("%s\n", pokemon.Pokemon.Name)
    }
}