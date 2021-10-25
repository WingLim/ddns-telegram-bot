package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

var bot, _ = tb.NewBot(tb.Settings{
	Token: os.Getenv("TOKEN"),
})

func init() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "https://ddns-bot.vercel.app"
	}

	bot.Handle("/start", func(m *tb.Message) {
		_, _ = bot.Send(m.Sender, "Welcome to use DDNS Bot!")
	})

	bot.Handle("/gethook", func(m *tb.Message) {
		chatId := m.Chat.ID
		menu := &tb.ReplyMarkup{}
		usage := menu.URL("Usage", "https://github.com/WingLim/ddns-telegram-bot/blob/main/README.md")
		menu.Inline(
			menu.Row(usage),
		)

		_, _ = bot.Send(m.Sender, fmt.Sprintf("Your Webhook URL:\n%s/api/hook?chatId=%d", host, chatId), menu)
	})
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tb.Update

	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error to parse Telegram request", err)
	}

	bot.ProcessUpdate(update)
}
