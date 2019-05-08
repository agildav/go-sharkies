package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// // // // // // // // // // // // // // // // // // // // // //

// Init gets all the environment variables
func Init() map[string]string {
	env := make(map[string]string)
	const envProduction string = "production"

	// Don't read .env files when deployed to Heroku for example
	//! Make sure the APP_ENV is present there
	if os.Getenv("APP_ENV") != envProduction {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("error loading .env file -> ", err)
		}
	}

	/*
		!: Add new env variables here
	*/

	// # APP Config
	env["APP_ENV"] = os.Getenv("APP_ENV")
	env["PORT"] = os.Getenv("PORT")

	// DB Config
	env["DB_USER"] = os.Getenv("DB_USER")
	env["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	env["DB_HOST"] = os.Getenv("DB_HOST")
	env["DB_PORT"] = os.Getenv("DB_PORT")
	env["DB_NAME"] = os.Getenv("DB_NAME")

	return env
}
