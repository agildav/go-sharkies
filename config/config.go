package config

import (
	"github.com/joho/godotenv"
	"log"
)

// // // // // // // // // // // // // // // // // // // // // //

// Init gets all the environment variables
func Init() map[string]string {
	var env map[string]string
	env, err := godotenv.Read()

	if err != nil {
		log.Fatal("error loading .env file -> ", err)
	}

	return env
}
