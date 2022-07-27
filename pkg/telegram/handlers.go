package telegram

import (
	"TestProjectForBtcTgBot/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const commandStart = "start"
const commandAddWallet = "add_wallet"
const commandRemoveWallet = "remove_wallet"
const commandLogin = "login"

func (b *Bot) handleCommand(message *tgbotapi.Message) {

	if message.Command() == commandStart {
		b.handleStartCommand(message)
		return
	}

	if message.Command() == commandLogin {
		b.handleLoginCommand(message)
		return
	}

	// check if user is logged in
	if models.IsUserLoggedIn(message.Chat.ID) == false {
		b.sendUserShouldBeLoggedIn(message)
		return
	}

	switch message.Command() {

	case commandAddWallet:
		resp := "Add wallet"
		b.SendMessage(resp, message.Chat.ID)
	case commandRemoveWallet:
		resp := "Remove wallet"
		b.SendMessage(resp, message.Chat.ID)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) {
	b.SendMessage("Unknown command", message.Chat.ID)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) {

	b.SendMessage("Hello!", message.Chat.ID)
	b.sendUserShouldBeLoggedIn(message)
}

func (b *Bot) handleLoginCommand(message *tgbotapi.Message) {
	m := strings.Split(message.Text, " ")

	login := m[1]
	password := m[2]

	user := models.LoadUserByLogin(login)

	if user.ID == 0 {
		b.sendLoginOrPasswordIsIncorrect(message)
		return
	}

	if !user.IsValidPasswordHash(password) {
		b.sendLoginOrPasswordIsIncorrect(message)
		return
	}

	b.authUser(login, message.Chat.ID)
	b.SendMessage("You are logged in, you can use bot functions", message.Chat.ID)
}

func (b *Bot) sendLoginOrPasswordIsIncorrect(message *tgbotapi.Message) {
	b.SendMessage("Login or password is incorrect", message.Chat.ID)
}

func (b *Bot) sendUserShouldBeLoggedIn(message *tgbotapi.Message) {
	b.SendMessage("login first using: /login <username> <password>", message.Chat.ID)
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	// check if user is logged in
	if models.IsUserLoggedIn(message.Chat.ID) == false {
		b.sendUserShouldBeLoggedIn(message)
		return
	}

	b.ReplyMessage(message.Text, message.Chat.ID, message.MessageID)
}
