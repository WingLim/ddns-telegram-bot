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
	Msg         string      `json:"text"`
	ChatID      int64       `json:"chat_id"`
	Method      string      `json:"method"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
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
	resp := BotResponse{
		ChatID: chatId,
		Method: "sendMessage",
	}
	text := ""
	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			text = "Welcome to use DDNS Bot!"
		case "gethook":
			text = fmt.Sprintf("Your Webhook URL:\n%s/api/hook?chatId=%d", host, chatId)
			inlineButtons := []tgbotapi.InlineKeyboardButton{}
			usage := tgbotapi.NewInlineKeyboardButtonURL("Usage", "https://github.com/WingLim/ddns-telegram-bot/blob/main/README.md")
			inlineButtons = append(inlineButtons, usage)
			inlineKM := tgbotapi.NewInlineKeyboardMarkup(inlineButtons)
			resp.ReplyMarkup = inlineKM
		}

		resp.Msg = text

		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}
