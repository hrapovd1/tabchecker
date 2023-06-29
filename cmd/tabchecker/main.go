package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

/*
1. read config:
{
	"left": "dsn",
	"right": "dsn",
	"table": {
		"name": "name",
		"fields": {
			"colonName1": "type",
			"colonName2": "type"
		}
	}
}
2. check colons for name and types
3. check count rows
4. check rows by index
*/

var ErrNoConfig = errors.New("Config not provided")

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

// return config path from args or error
func getConfig() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", ErrNoConfig
	}
	return filepath.Abs(args[1])
}
