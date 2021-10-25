package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

type IP struct {
	Result  string `json:"result"`
	Addr    string `json:"addr"`
	Domains string `json:"domains"`
}

type DDNSRequest struct {
	IPv4 IP `json:"ipv4,omitempty"`
	IPv6 IP `json:"ipv6,omitempty"`
}

type HookResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg,omitempty"`
}

const Template = "%s: %s\nIP: %s\nDomains: %s\n"

func HookHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	var req DDNSRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Fatal("Error to parse DDNS request: ", err)
	}

	chatId, _ := strconv.ParseInt(r.URL.Query().Get("chatId"), 10, 64)

	text := ""
	if req.IPv4.Addr != "" {
		text += fmt.Sprintf(Template, "IPv4", req.IPv4.Result, req.IPv4.Addr, req.IPv4.Domains)
	}

	if req.IPv6.Addr != "" {
		text += fmt.Sprintf(Template, "IPv6", req.IPv6.Result, req.IPv6.Addr, req.IPv6.Domains)
	}

	chat := &tb.Chat{
		ID: chatId,
	}

	var resp HookResponse

	if _, err := bot.Send(chat, text); err != nil {
		resp = HookResponse{
			Status: false,
			Msg:    err.Error(),
		}
	} else {
		resp = HookResponse{
			Status: true,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
