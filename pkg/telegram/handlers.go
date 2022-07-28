package telegram

import (
	"TestProjectForBtcTgBot/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const commandStart = "start"
const commandWalletsList = "wallets_list"
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
		b.handleAddWallet(message)
	case commandRemoveWallet:
		b.handleDeleteWallet(message)
	case commandWalletsList:
		b.handleWalletsList(message)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleWalletsList(message *tgbotapi.Message) {
	user := models.GetUserByChatId(message.Chat.ID)
	wallets := models.GetAllWalletsByUserID(user.ID)

	b.sendWalletsList(wallets, message)
}

func (b *Bot) handleAddWallet(message *tgbotapi.Message) {
	cmd := strings.Split(message.Text, " ")

	if len(cmd) < 2 {
		b.sendWrongCommandUsage(message)
		return
	}

	walletAddr := cmd[1]

	user := models.GetUserByChatId(message.Chat.ID)
	wallet := models.NewWallet()

	wallet.Wallet = walletAddr
	wallet.UserID = user.ID

	if _, err := models.AddWallet(wallet); err != nil {
		b.sendWalletWasNotAdded(message)
		return
	}

	b.sendWalletWasAdded(wallet.Wallet, message)
}

func (b *Bot) handleDeleteWallet(message *tgbotapi.Message) {
	cmd := strings.Split(message.Text, " ")
	if len(cmd) < 2 {
		b.sendWrongCommandUsage(message)
		return
	}
	walletAddr := cmd[1]
	user := models.GetUserByChatId(message.Chat.ID)
	wallet := models.NewWallet()
	wallet.Wallet = walletAddr
	wallet.UserID = user.ID

	if err := models.DeleteWallet(wallet); err != nil {
		b.sendWalletWasNotDeleted(message)
		return
	}

	b.sendWalletWasDeleted(message)
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

	if models.IsUserLoggedIn(message.Chat.ID) == true {
		b.SendMessage("You are logged in already", message.Chat.ID)
		return
	}

	if len(m) < 2 {
		b.SendMessage("Wrong command usage enter login and password please (/login <username> <password>)", message.Chat.ID)
		return
	}

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

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	// check if user is logged in
	if models.IsUserLoggedIn(message.Chat.ID) == false {
		b.sendUserShouldBeLoggedIn(message)
		return
	}

	b.ReplyMessage(message.Text, message.Chat.ID, message.MessageID)
}
