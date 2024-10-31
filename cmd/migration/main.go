package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"shortener-smile/database"
	_ "shortener-smile/migration"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "/app/migration/", "directory with migration files")
)

func main() {
	err := flags.Parse(os.Args[1:])

	if err != nil {
		fmt.Println(err)
		return
	}
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}

	fmt.Println(fmt.Sprintf("Migrating from %s", *dir))

	command := args[1]

	db, err := database.GetConnection()
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	if err := goose.RunContext(context.Background(), command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
