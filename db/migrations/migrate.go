package main

import (
	"flag"
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/db"
	models2 "github.com/elizielx/arcturus-api/internal/models"
	"github.com/elizielx/arcturus-api/internal/utils"
	"log"
	"math/rand"
	"time"
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

		err := seedUsers()
		if err != nil {
			log.Fatal(err)
			return
		}

		err = seedPolls()
		if err != nil {
			log.Fatal(err)
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

func seedPolls() error {
	polls := []models2.Poll{
		{Title: "Qual a melhor linguagem de programação? 1", Description: "Qual a melhor linguagem de programação?", Deadline: time.Now().Add(time.Hour * 2), CreatedBy: uint64(generateNumberBetween(1, 2))},
		{Title: "Qual a melhor linguagem de programação? 2", Description: "Qual a melhor linguagem de programação?", Deadline: time.Now().Add(time.Hour * 2), CreatedBy: uint64(generateNumberBetween(1, 2))},
		{Title: "Qual a melhor linguagem de programação? 3", Description: "Qual a melhor linguagem de programação?", Deadline: time.Now().Add(time.Hour * 2), CreatedBy: uint64(generateNumberBetween(1, 2))},
		{Title: "Qual a melhor linguagem de programação? 4", Description: "Qual a melhor linguagem de programação?", Deadline: time.Now().Add(time.Hour * 2), CreatedBy: uint64(generateNumberBetween(1, 2))},
	}

	for i := range polls {
		if err := db.GetDatabase().Create(&polls[i]).Error; err != nil {
			log.Fatal(err)
			return err
		}
	}

	log.Println("Polls successfully seeded")
	return nil
}

func generateNumberBetween(min, max int) int {
	return min + rand.Intn(max-min)
}
