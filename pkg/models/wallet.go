package models

import (
	"TestProjectForBtcTgBot/pkg/database"
	"database/sql"
	"errors"
	"log"
	"time"
)

type Wallet struct {
	ID              int    `db:"wallets.id"`
	UserID          int    `db:"wallets.user_id"`
	Wallet          string `db:"wallets.wallet"`
	LastTransaction int64  `db:"wallets.last_transaction"`
}

func NewWallet() Wallet {
	return Wallet{}
}

func AddWallet(wallet Wallet) (Wallet, error) {
	if walletExists(wallet) {
		return wallet, errors.New("wallet exists")
	}

	lastT := time.Now().Unix()

	_, err := database.Instance().Exec("INSERT INTO wallets (user_id, wallet, last_transaction) values (?, ?, ?)", wallet.UserID, wallet.Wallet, lastT)
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

func GetAllWallets() []Wallet {
	rows, _ := database.Instance().Query("SELECT wallet, user_id, id, last_transaction FROM wallets")
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

func UpdateWalletLastTransaction(wallet Wallet, lastTransaction int64) error {
	_, err := database.Instance().Exec("UPDATE wallets set last_transaction = ? WHERE id = ?", lastTransaction, wallet.ID)
	if err != nil {
		return err
	}

	return nil
}
