package btc

import (
	"TestProjectForBtcTgBot/pkg/models"
	"TestProjectForBtcTgBot/pkg/telegram"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type HistoryResponse struct {
	History []HistoryItem
}

type HistoryItem struct {
	Time  int64    `json:"time"`
	Addr  []string `json:"addr"`
	Value int64    `json:"value"`
	Txid  string   `json:"txid"`
}

func GetNewTransactionsByWallet(wallet models.Wallet) ([]HistoryItem, error) {

	payloadString := fmt.Sprintf(`{
    	"addr": "%s"
	}`, wallet.Wallet)

	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.blockonomics.co/api/searchhistory", payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer sHmwJJBu1z5qR8Pz6RzJ7odOXlFxz06QAyIwC8cz63U")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var history HistoryResponse
	json.Unmarshal(body, &history)
	fmt.Println(history)

	lastCheck := wallet.LastTransaction

	var newTransactions []HistoryItem

	for _, history := range history.History {
		if time.Unix(history.Time, 0).After(time.Unix(lastCheck, 0)) {
			newTransactions = append(newTransactions, history)
		}
	}

	return newTransactions, nil
}

func CheckAllWalletsAndSendNotifications(bot *telegram.Bot) {
	wallets := models.GetAllWallets()

	for _, wallet := range wallets {
		newlyCreatedTransactions, _ := GetNewTransactionsByWallet(wallet)

		for _, tr := range newlyCreatedTransactions {

			//if tr.Value > 0 {
			user := models.GetUserById(wallet.UserID)

			fmt.Printf("value %d \n", tr.Value)
			bot.SendIncomingTransactionNotification(tr.Value, user.ChatId, wallet)
			//}
		}

		if len(newlyCreatedTransactions) > 0 {
			if err := models.UpdateWalletLastTransaction(wallet, time.Now().Unix()); err != nil {
				log.Fatal(err)
			}
		}

	}
}
