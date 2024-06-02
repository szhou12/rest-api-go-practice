package main

import (
	"fmt"
	"os"
)

/*
Respnsible for having a hold of the env variables in the run time of the server from the injected variables that we're going to have
Usually the best practice to have the env vars injected in the runtime instead of hardcoding them
*/
type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string // used for authentication
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "projectmanager"),
		JWTSecret: getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}

/*
key: the key name of the env variable in env file
fallback: the default value if the env variable is not found
*/
func getEnv(key, fallback string) string {
	// trick
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
