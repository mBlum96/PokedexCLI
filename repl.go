package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct{
    name string
    description string
    callback func() error
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

func handleCommand(command string) error{
    cmd,ok := commands[command]
    if(!ok){
        var err_str string = fmt.Sprintf("command '%s' is unknown", command)
        return errors.New(err_str)
    }
    cmd.callback()
    return nil
}


func repl() {
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