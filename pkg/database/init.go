package database

import (
	"TestProjectForBtcTgBot/pkg/hash"
	"fmt"
)

const createUser string = `
  CREATE TABLE IF NOT EXISTS users (
	  id INTEGER PRIMARY KEY,
	  login TEXT NOT NULL,
	  password_hash TEXT NOT NULL,
	  chat_id INTEGER DEFAULT NULL,
	  is_logged_in INTEGER DEFAULT 0
  );`

const createWallets string = `
	CREATE TABLE IF NOT EXISTS wallets (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER NOT NULL,
	  wallet TEXT NOT NULL UNIQUE,
	  last_transaction TEXT
);`

func InitDB() error {
	if _, err := dbInstance.Exec(createUser); err != nil {
		return err
	}

	if _, err := dbInstance.Exec(createWallets); err != nil {
		return err
	}

	pwd, _ := hash.HashPassword("qwerty")

	insertQuery := fmt.Sprintf("INSERT INTO users (login, password_hash) VALUES ('username', '%s')", pwd)

	dbInstance.Exec(insertQuery)

	return nil
}
