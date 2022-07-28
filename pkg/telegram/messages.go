package telegram

import (
	"TestProjectForBtcTgBot/pkg/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendLoginOrPasswordIsIncorrect(message *tgbotapi.Message) {
	b.SendMessage("Login or password is incorrect", message.Chat.ID)
}

func (b *Bot) sendUserShouldBeLoggedIn(message *tgbotapi.Message) {
	b.SendMessage("Hello!", message.Chat.ID)
	b.SendMessage("login first using: /login <username> <password>", message.Chat.ID)
}

func (b *Bot) sendWrongCommandUsage(message *tgbotapi.Message) {
	b.SendMessage("Wrong command usage", message.Chat.ID)
}

func (b *Bot) sendWalletWasNotAdded(message *tgbotapi.Message) {
	b.SendMessage("Wallet was not added", message.Chat.ID)
}

func (b *Bot) sendWalletWasAdded(walletAddr string, message *tgbotapi.Message) {
	b.SendMessage(fmt.Sprintf("Wallet '%s' was added", walletAddr), message.Chat.ID)
}

func (b *Bot) sendWalletWasDeleted(message *tgbotapi.Message) {
	b.SendMessage("Wallet was not deleted", message.Chat.ID)
}

func (b *Bot) sendWalletWasNotDeleted(message *tgbotapi.Message) {
	b.SendMessage("Wallet was deleted", message.Chat.ID)
}

func (b *Bot) sendWalletsList(walletsList []models.Wallet, message *tgbotapi.Message) {
	text := "Your wallets list:"

	fmt.Println(walletsList)

	for _, w := range walletsList {
		text += fmt.Sprintf("\n%s", w.Wallet)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	b.bot.Send(msg)
}
