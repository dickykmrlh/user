package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dickykmrlh/profile/config"
)

var db *sql.DB

func NewDBConn() {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.GetConfig().DB.Host,
		config.GetConfig().DB.Port,
		config.GetConfig().DB.DBName,
		config.GetConfig().DB.Username,
		config.GetConfig().DB.Password,
		config.GetConfig().DB.SSLMode)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
