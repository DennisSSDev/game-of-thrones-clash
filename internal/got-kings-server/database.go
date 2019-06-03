package server

import (

	//driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = []string{
	`CREATE TABLE IF NOT EXISTS king (
		full_name TEXT NOT NULL,
		house TEXT NOT NULL,
		king_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
	);`,
	`CREATE TABLE IF NOT EXISTS conqueredland (
    land TEXT NOT NULL,
    king_id INT UNSIGNED NOT NULL,
	land_id INT UNSIGNED NOT NULL,
	PRIMARY KEY(king_id, land_id),
	FOREIGN KEY(king_id) REFERENCES king(king_id)
	);`,
	`CREATE TABLE IF NOT EXISTS alliance (
	king_id INT UNSIGNED NOT NULL,
	alliance TEXT NOT NULL,
	alliance_id INT UNSIGNED NOT NULL,
	PRIMARY KEY(king_id, alliance_id)
)`,
}

// DatabaseHandle attempt to abstract all the low level functionality of dealing with the database connection
type DatabaseHandle struct {
	DB *sqlx.DB
}

// Database holds all the required functionality for the handler
type Database interface {
	Connect(dbName string, dbAddress string) error
}

// Connect attempts to connect to the providecd dbName and dbAddress
func (db *DatabaseHandle) Connect(dbAddress string) error {
	database := sqlx.MustConnect("mysql", dbAddress)
	db.DB = database
	for _, value := range schema {
		db.DB.MustExec(value)
	}
	return nil
}
