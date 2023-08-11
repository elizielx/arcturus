package db

import (
	"fmt"
	"github.com/elizielx/arcturus-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDatabase(configuration config.Configuration) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", configuration.DatabaseHost, configuration.DatabaseUser, configuration.DatabasePassword, configuration.DatabaseName, configuration.DatabasePort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	checkRoleLevel := db.Exec("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_level');")
	if checkRoleLevel.RowsAffected == 0 {
		rolesLevel := db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_level') THEN CREATE TYPE role_level AS ENUM ('USER', 'ADMIN'); END IF; END $$;")
		if rolesLevel.Error != nil {
			log.Fatal(rolesLevel.Error)
		}
		log.Println("Role level successfully created")
	}

	log.Println("Database successfully connected")
}

func GetDatabase() *gorm.DB {
	return db
}
