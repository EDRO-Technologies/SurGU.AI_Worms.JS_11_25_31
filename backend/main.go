package main

import (
	"log"
	"worm-pack/cmd"
	"worm-pack/config"

	"worm-pack/db"

	"github.com/joho/godotenv"
)

func main() {

	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	confVars, configErr := config.New()

	examples := db.Init_db()

	_ = examples

	if configErr != nil {
		log.Fatal(configErr)
	}

	app := cmd.InitApp()

	err := app.Listen(confVars.Port)
	if err != nil {
		log.Fatal(err)
	}
}
