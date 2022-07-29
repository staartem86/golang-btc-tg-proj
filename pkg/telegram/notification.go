package telegram

import (
	"TestProjectForBtcTgBot/pkg/models"
	"fmt"
)

func (b *Bot) SendIncomingTransactionNotification(value int64, chatID int64, wallet models.Wallet) {
	btcVal := float64(value) / 100000000

	b.SendMessage(fmt.Sprintf("You received %.8f btc on: %s \n", btcVal, wallet.Wallet), chatID)
}
