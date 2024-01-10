package main

import (
	"fmt"
	"sync"
)

type response struct{
    Results []LocationData `json:"results"`
    Next string `json:"next"`
    Previous string `json:"previous"`
}

type LocationData struct {
    Name string `json:"name"`
}

var lock sync.Mutex
var result response
var currentMapPage string = "https://pokeapi.co/api/v2/location/"

func mapPrinter(result *response){
    for _,area := range(result.Results){
        fmt.Printf("%s\n", area.Name)
    }
}