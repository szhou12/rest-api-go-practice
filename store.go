package main

import "database/sql"

/*
Define Store as interface so that we can easily inject it into services and because of that, we can easily test them bc we can mock this interface
*/ 
type Store interface {
	// Users
	CreateUser() error
}

/*
Responsible for communicating with the database
*/
type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser() error {
	return nil
}