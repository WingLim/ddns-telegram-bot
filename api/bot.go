package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotResponse struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
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
			text = fmt.Sprintf("Your Webhook URL:\nhttps://ddns-bot.vercel.dev/api/hook?chatId=%d", chatId)
		}

		data := BotResponse{
			Msg:    text,
			ChatID: chatId,
			Method: "sendMessage",
		}

		msg, _ := json.Marshal(data)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(msg))
	}
}
