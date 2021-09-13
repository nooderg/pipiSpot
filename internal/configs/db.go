package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func init() {
	loadDBClient()
}

func loadDBClient() error {
	sqlCred := fmt.Sprintf("%s:%s@tcp(%s:%s)", os.Getenv("dbuser"), os.Getenv("dbpassword"), os.Getenv("dburl"), os.Getenv("dbport"))

	for {
		log.Println("Connecting to DB")
		database, err := sql.Open("mysql", sqlCred)
		if err != nil {
			log.Println("Connection failled, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		DBClient = database
		break
	}

	log.Println("Connexion to DB successful")
	return nil
}


func GetDBClient() *sql.DB {
	if DBClient.Ping() != nil {
		loadDBClient()
		return DBClient
	} else {
		return DBClient
	}
}