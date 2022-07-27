package telegram

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	db  *sql.DB
}

func NewBot(apiKey string, db *sql.DB) *Bot {
	tgBot, err := tgbotapi.NewBotAPI(apiKey)
	tgBot.Debug = true

	if err != nil {
		log.Panic(err)
	}

	return &Bot{bot: tgBot, db: db}
}

func (b *Bot) SendMessage(text string, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, text)
	b.bot.Send(msg)
}

func (b *Bot) ReplyMessage(text string, chatId int64, msgId int) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyToMessageID = msgId
	b.bot.Send(msg)
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}
