package main

import (
	"flag"
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/db"
	models2 "github.com/elizielx/arcturus-api/internal/models"
	"github.com/elizielx/arcturus-api/internal/utils"
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
	seedFlag := flag.Bool("seed", false, "Seed database")
	flag.Parse()

	if *resetFlag {
		log.Println("Dropping database")
		err := db.GetDatabase().Migrator().DropTable(&models2.User{}, &models2.Division{}, &models2.Poll{}, &models2.Choice{}, &models2.Vote{})
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	err := db.GetDatabase().AutoMigrate(&models2.User{}, &models2.Division{}, &models2.Poll{}, &models2.Choice{}, &models2.Vote{})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Migration successfully")

	if *seedFlag {
		log.Println("Seeding database")
		users := seedUsers()
		if users != nil {
			log.Fatal(users)
			return
		}
	}

}

func seedUsers() error {
	users := []models2.User{
		{Username: "admin", Password: "admin", Role: models2.ADMIN},
		{Username: "user", Password: "user", Role: models2.USER},
	}

	for i := range users {
		hashedPassword, err := utils.HashPassword(users[i].Password)
		if err != nil {
			log.Fatal(err)
			return err
		}
		users[i].Password = hashedPassword

		if err := db.GetDatabase().Create(&users[i]).Error; err != nil {
			log.Fatal(err)
			return err
		}
	}

	log.Println("Users successfully seeded")
	return nil
}
