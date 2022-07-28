package main

import (
	"TestProjectForBtcTgBot/pkg/database"
	"TestProjectForBtcTgBot/pkg/telegram"
	dotenv "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {

	if err := dotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db := database.Connect()
	defer db.Close()

	database.InitDB()

	bot := telegram.NewBot(os.Getenv("TG_BOT_API_KEY"), db)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}

}
