package models

import (
	"TestProjectForBtcTgBot/pkg/database"
	"TestProjectForBtcTgBot/pkg/hash"
	"log"
)

type User struct {
	ID           int    `db:"users.id"`
	PasswordHash string `db:"users.password_hash"`
	Login        string `db:"users.login"`
	ChatId       int64  `db:"users.chat_id"`
	IsLoggedIn   int    `db:"users.is_logged_in"`
}

func GetUserByChatId(chatId int64) User {
	row := database.Instance().QueryRow("SELECT * FROM users WHERE chat_id = ?", chatId)
	if row == nil {
		log.Fatal(row)
	}

	var user User

	row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.ChatId, &user.IsLoggedIn)

	return user
}

func GetUserById(id int) User {
	row := database.Instance().QueryRow("SELECT * FROM users WHERE id = ?", id)
	if row == nil {
		log.Fatal(row)
	}

	var user User

	row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.ChatId, &user.IsLoggedIn)

	return user
}

func IsUserLoggedIn(tgChatId int64) bool {
	user := GetUserByChatId(tgChatId)
	if user.ID == 0 {
		return false
	}

	if user.IsLoggedIn == 0 {
		return false
	}

	return true
}

func LoadUserByLogin(login string) User {
	row := database.Instance().QueryRow("SELECT * FROM users WHERE login = ?", login)
	if row == nil {
		log.Fatal(row)
	}

	var user User

	row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.ChatId, &user.IsLoggedIn)

	return user
}

func (u *User) IsValidPasswordHash(password string) bool {
	return hash.CompareHashedPassword(u.PasswordHash, password)
}

func AuthUser(login string, chatID int64) {
	_, err := database.Instance().Exec("UPDATE users SET chat_id = ?, is_logged_in = 1 WHERE login = ?", chatID, login)
	if err != nil {
		log.Fatal(err)
	}
}
