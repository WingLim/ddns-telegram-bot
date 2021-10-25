package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotResponse struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("HOST")
	if host == "" {
		host = "https://ddns-bot.vercel.app"
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error to parse Telegram request", err)
	}

	chatId := update.Message.Chat.ID
	text := ""
	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			text = "Welcome to use DDNS Bot!"
		case "gethook":
			text = fmt.Sprintf("Your Webhook URL:\n%s/api/hook?chatId=%d", host, chatId)
		}

		resp := BotResponse{
			Msg:    text,
			ChatID: chatId,
			Method: "sendMessage",
		}

		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}
