package main

import (
	"TestProjectForBtcTgBot/pkg/btc"
	"TestProjectForBtcTgBot/pkg/database"
	"TestProjectForBtcTgBot/pkg/telegram"
	"fmt"
	dotenv "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	tasks "github.com/theTardigrade/golang-tasks"
	"log"
	"os"
	"time"
)

func main() {

	if err := dotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db := database.Connect()
	defer db.Close()

	database.InitDB()

	bot := telegram.NewBot(os.Getenv("TG_BOT_API_KEY"), db)

	// task for checking wallets every minute
	tasks.Set(time.Minute/10, false, func(id *tasks.Identifier) {
		fmt.Println("Checking wallets")
		btc.CheckAllWalletsAndSendNotifications(bot)
	})

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}

}
