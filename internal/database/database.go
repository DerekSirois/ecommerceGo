package database

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func InitDb() {
	dsn := os.Getenv("DSN")

	// 5 retries
	for i := 0; i < 5; i++ {
		db, err := sqlx.Connect("postgres", dsn)
		if err == nil && db != nil {
			Db = db
			log.Println("Connected to the database")
			return
		}
		log.Println("Error... retrying in 2 seconds")
		time.Sleep(time.Second * 2)
	}
	log.Fatal("Couln't initialize the database...")
}
