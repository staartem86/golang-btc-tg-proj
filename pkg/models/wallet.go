package models

import (
	"TestProjectForBtcTgBot/pkg/database"
	"database/sql"
	"errors"
	"log"
)

type Wallet struct {
	ID              int            `db:"wallets.id"`
	UserID          int            `db:"wallets.user_id"`
	Wallet          string         `db:"wallets.wallet"`
	LastTransaction sql.NullString `db:"wallets.last_transaction"`
}

func NewWallet() Wallet {
	return Wallet{}
}

func AddWallet(wallet Wallet) (Wallet, error) {
	if walletExists(wallet) {
		return wallet, errors.New("wallet exists")
	}

	_, err := database.Instance().Exec("INSERT INTO wallets (user_id, wallet) values (?, ?)", wallet.UserID, wallet.Wallet)
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func DeleteWallet(wallet Wallet) error {
	if walletExists(wallet) {
		return errors.New("wallet not deleted")
	}

	_, err := database.Instance().Exec("DELETE FROM wallets WHERE wallet = ?", wallet.Wallet)
	if err != nil {
		return err
	}

	return nil
}

func GetAllWalletsByUserID(userID int) []Wallet {
	rows, _ := database.Instance().Query("SELECT wallet, user_id, id, last_transaction FROM wallets WHERE user_id = ?", userID)
	var wallets []Wallet

	for rows.Next() {
		w := Wallet{}
		if err := rows.Scan(&w.Wallet, &w.UserID, &w.ID, &w.LastTransaction); err != nil {
			log.Fatal(err)
		}
		wallets = append(wallets, w)
	}

	return wallets
}

func walletExists(wallet Wallet) bool {
	var exists bool
	err := database.Instance().QueryRow("SELECT exists (SELECT id FROM wallets WHERE wallet = ?)", wallet.Wallet).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	return exists
}
