package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/go-services/schema"
	"github.com/jacky-htg/go-services/services/config"
	"github.com/jacky-htg/go-services/services/database"
)

func main() {

	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	flag.Parse()

	// =========================================================================
	// Start Database

	dbx, err := database.Openx()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer dbx.Close()

	switch flag.Arg(0) {
	case "migrate":
		db, err := database.Open()
		if err != nil {
			log.Fatalf("error: connecting to db: %s", err)
		}
		defer db.Close()
		if err := schema.Migrate(db); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return

	case "seed":
		if err := schema.Seed(dbx); err != nil {
			log.Println("error seeding database", err)
			os.Exit(1)
		}
		log.Println("Seed data complete")
		return
	}
}
