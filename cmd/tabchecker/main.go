package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

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
2.a get namespaces, tables and compare
3. check count rows
4. check rows by index
*/

func main() {
	// read config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("config = %v\n", &config)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	db, err := sql.Open("mysql", config.Left)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "show tables;")
	if err != nil {
		log.Fatalln(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalln(err)
	}

	values := make([]any, len(columns))
	for i := range values {
		values[i] = new([]byte)
	}

	for _, colonName := range columns {
		fmt.Printf("  %v  ", colonName)
	}
	fmt.Print("\n")
	fmt.Println("-----------------------")
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			log.Fatalln(err)
		}
		for _, val := range values {
			fmt.Printf("%s	", val.(string))
		}
		fmt.Print("\n")
	}
}
