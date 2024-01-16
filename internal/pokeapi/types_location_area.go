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
	Pokemon []struct {
		PokemonDetails struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	}
}
