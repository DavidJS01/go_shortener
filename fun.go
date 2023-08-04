package main

import (
	"errors"
	db "example.com/mod/db"
	"fmt"
	"os"
)

func parseArgs() ([]string, error) {
	args := os.Args
	if len(args) < 2 {
		fmt.Print("Incorrect length of arguments provided")
		return []string{""}, errors.New("incorrect length of arguments provided")
	}
	return args, nil
}

func main() {
	args, err := parseArgs()
	if err != nil {
		return
	}
	if args[2] == "true" {
		mapping, _ := db.GetUrlMapping(args[1], true)
		fmt.Print(mapping)
	} else {
		mapping, _ := db.GetUrlMapping(args[1], false)
		fmt.Print(mapping)
	}

}
