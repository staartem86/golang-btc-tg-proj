package telegram

import "TestProjectForBtcTgBot/pkg/models"

func (b *Bot) authUser(login string, chatId int64) {
	models.AuthUser(login, chatId)
}
