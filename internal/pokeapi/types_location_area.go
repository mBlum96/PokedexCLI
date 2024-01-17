package pokeapi

type LocationResponse struct{
    Locations []locationData `json:"results"`
    Next string `json:"next"`
    Previous string `json:"previous"`
}

type locationData struct {
    Name string `json:"name"`
}

type LocationAddresses struct{
	Current string
	Next string
	Previous string
}

type pokemonEncountered struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}