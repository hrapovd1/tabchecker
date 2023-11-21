package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"

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

	leftDb, err := sql.Open("mysql", config.Left.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer leftDb.Close()

	if err := getTables(ctx, leftDb, config.Left.Type); err != nil {
		log.Fatalln(err)
	}

	rightDb, err := sql.Open("pgx", fmt.Sprint(config.Right.Type+"://"+config.Right.DSN))
	if err != nil {
		log.Fatalln(err)
	}
	defer rightDb.Close()
	if err := getTables(ctx, rightDb, config.Right.Type); err != nil {
		log.Fatalln(err)
	}
}

func getTables(ctx context.Context, db *sql.DB, typeDb string) error {
	query := ""
	switch typeDb {
	case "mysql":
		query = "show tables;"
	case "postgres":
		query = "SELECT FORMAT('%s.%s', schemaname,tablename) tables FROM pg_catalog.pg_tables;"
	default:
	}
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
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
			return err
		}
		for _, val := range values {
			valResult := val.(*[]uint8)
			fmt.Printf("%s	", string(*valResult))
		}
		fmt.Print("\n")
	}

	return nil
}
