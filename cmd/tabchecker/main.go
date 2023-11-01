package main

import (
	"fmt"
	"log"

	"github.com/hrapovd1/tabchecker/internal/config"
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

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("config = %v\n", &config)
}
