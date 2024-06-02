package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

/*
This is going to be the implementation for MySQL and it's going to receive the connection to the database
*/
type MySQLStorage struct {
	db *sql.DB
}

/*
Here: initialize the database and make the connection
*/
func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// ensure successfully created a connection to the database by using Ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Connected to MySQL!")

	return &MySQLStorage{db: db}
}

/*
Return the actual pointer to the database and an error
Reason for having this method: We need to initialize the tables for the database and then we just return the databse
*/
func (s *MySQLStorage) Init() (*sql.DB, error) {
	// Initialize the tables 


	
	return s.db, nil
}