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
	seedFlag := flag.Bool("seed", false, "Seed database")
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
	users := []models.User{
		{Username: "admin", Password: "admin", Role: models.ADMIN},
		{Username: "user", Password: "user", Role: models.USER},
	}

	for _, user := range users {
		if err := db.GetDatabase().Create(&user).Error; err != nil {
			log.Fatal(err)
			return err
		}
	}

	log.Println("Users successfully seeded")
	return nil
}
