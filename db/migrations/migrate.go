package main

import (
	"flag"
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/db"
	"github.com/elizielx/arcturus-api/models"
	"log"
)

func init() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		log.Fatal(err)
	}

	db.InitDatabase(configuration)
}

func main() {
	resetFlag := flag.Bool("reset", false, "Reset database")
	flag.Parse()

	if *resetFlag {
		log.Println("Dropping database")
		err := db.GetDatabase().Migrator().DropTable(&models.User{}, &models.Division{}, &models.Poll{}, &models.Choice{}, &models.Vote{})
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	err := db.GetDatabase().AutoMigrate(&models.User{}, &models.Division{}, &models.Poll{}, &models.Choice{}, &models.Vote{})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Migration successfully")
}
