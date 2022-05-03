package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nooderg/pipiSpot/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func init() {
	loadDBClient(true)
}

func loadDBClient(fillDB bool) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("dbhost"),
		os.Getenv("dbuser"),
		os.Getenv("dbpassword"),
		os.Getenv("dbname"),
		os.Getenv("dbport"),
	)

	for {
		log.Println("Connecting to DB...")
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			log.Println("Connection failled, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		if fillDB {
			for {
				err = db.AutoMigrate(&models.User{}, &models.Spot{}, &models.Comment{}, &models.Ratings{})
				if err != nil {
					log.Println("Migration failled, retrying in 5 seconds...")
					time.Sleep(5 * time.Second)
					continue
				}
				break
			}
		}
		
		DBClient = db
		break
	}

	log.Println("Connexion to DB successful")
	return nil
}

func GetDBClient() *gorm.DB {
	db, err := DBClient.DB()
	if err != nil {
		return nil
	}
	if db.Ping() != nil {
		loadDBClient(false)
		return DBClient
	} else {
		return DBClient
	}
}
