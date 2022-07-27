package main

import (
	"TestProjectForBtcTgBot/pkg/database"
	"TestProjectForBtcTgBot/pkg/telegram"
	dotenv "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	if err := dotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db := database.Connect()
	defer db.Close()

	database.InitDB()

	bot := telegram.NewBot("5419504507:AAGiOp6S5LE82KgBKk20q5Tkfh4ta5b0g0Y", db)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}

}
