package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dickykmrlh/user/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func NewDBConn() {
	log.Print("connecting to DB")
	dbConfig := config.GetConfig().GetDBConfig()

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.SSLMode)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(fmt.Errorf("opening db conn error: %s", err))
	}

	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetMaxOpenConns(dbConfig.MaxOpenCons)
	db.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifetime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(dbConfig.ConnMaxIdleTime))

	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Errorf("ping db conn error: %s", err))
	}

	log.Print("connected to DB")
}
