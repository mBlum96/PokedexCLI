package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type cliCommand struct{
    name string
    description string
    callback func() error
}

type LocationData struct {
    Name string `json:"name"`
}

type response struct{
    Results []LocationData `json:"results"`
    Next string `json:"next"`
    Previous string `json:"previous"`
}

var lock sync.Mutex

type mapPage struct{
    count int
}

func (p *mapPage) Increment(){
    p.count++
}

func (p *mapPage) Decrement(){
    p.count--
}

func (p *mapPage) GetCount() int{
    return p.count
}

var commands map[string]cliCommand

func init() {
    commands = map[string]cliCommand{
        "help": {
            name: "help",
            description: "Show help",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the program",
            callback: commandExit,
        },
        "map":{
            name: "map",
            description: "Show the map",
            callback: commandMap,
        },
        "bmap":{
            name: "bmap",
            description: "Show the previous map",
            callback: commandBMap,
        },
    }
}

var result response
var currentMapPage string = "https://pokeapi.co/api/v2/location/"

func mapPrinter(result *response){
    for _,area := range(result.Results){
        fmt.Printf("%s\n", area.Name)
    }
}

func commandMap() error{
    resp, err := http.Get(currentMapPage)
    if err != nil{
        return err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return errors.New("Error reading response body")
    }
    err = json.Unmarshal(body, &result)
    currentMapPage = result.Next
    mapPrinter(&result)
    return nil
}

func commandBMap() error{
    currentMapPage = result.Previous
    resp, err := http.Get(currentMapPage)
    if err != nil{
        return err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return errors.New("Error reading response body")
    }
    err = json.Unmarshal(body, &result)
    mapPrinter(&result)
    return nil
}

func commandHelp() error{
    fmt.Println("Available commands:")
    for _,cmd := range(commands){
        //print command name and description
        fmt.Printf("%v: %v\n",cmd.name,cmd.description)
    }
    return nil
}

func commandExit() error{
    fmt.Println("Exiting the pokedex...")
    os.Exit(0)
    return nil
}

func handleCommand(command string) error{
    cmd,ok := commands[command]
    if(!ok){
        var err_str string = fmt.Sprintf("command '%s' is unknown", command)
        return errors.New(err_str)
    }
    cmd.callback()
    return nil
}


func main(){
    reader := bufio.NewReader(os.Stdin)
    for{
        fmt.Print("pokedex>")
        input, err:= reader.ReadString('\n')
        if err != nil{
            fmt.Println("Error reading input:", err)
            continue
        }
        input = strings.TrimSpace(input)
        lock.Lock()
        err = handleCommand(input)
        lock.Unlock()
        if(err != nil){
            fmt.Println("Error executing command:", err)
            continue
        }
    }
}
