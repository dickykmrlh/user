package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dickykmrlh/user/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

func GetDB() *sql.DB {
	if dbInstance == nil {
		InitDBConn()
		return dbInstance
	}

	return dbInstance
}

func InitDBConn() {
	dbConfig := config.GetConfig().GetDBConfig()

	log.Print("connecting to DB")

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.SSLMode)

	var err error
	dbInstance, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("opening db conn error,", err)
	}

	dbInstance.SetMaxIdleConns(dbConfig.MaxIdleConns)
	dbInstance.SetMaxOpenConns(dbConfig.MaxOpenCons)
	dbInstance.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifetime))
	dbInstance.SetConnMaxIdleTime(time.Second * time.Duration(dbConfig.ConnMaxIdleTime))

	err = dbInstance.Ping()
	if err != nil {
		log.Fatal("ping db conn error, ", err)
	}
	log.Print("connected to DB")

	runMigration(dbInstance)
}

func runMigration(db *sql.DB) {
	dbConfig := config.GetConfig().GetDBConfig()

	if !dbConfig.MigrationRun {
		return
	}

	log.Print("migration run")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("migration failed, setup driver, ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", dbConfig.MigrationPath),
		"postgres", driver)
	if err != nil {
		log.Fatal("migration failed, setup migrate instance, ", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration failed, up, ", err)
	}
	log.Print("migration completed")
}
