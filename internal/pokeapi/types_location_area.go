package pokeapi

type LocationResponse struct{
    Locations []locationData `json:"results"`
    Next string `json:"next"`
    Previous string `json:"previous"`
}

type locationData struct {
    Name string `json:"name"`
}
