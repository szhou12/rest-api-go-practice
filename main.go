package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// MySQL storage configuration
	sqlStorage := NewMySQLStorage(cfg)

	// Initialize the database
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the store
	store := NewStore(db)

	// Inject the store into the API server
	api := NewAPIServer(":3000", store)
	api.Serve()
}
