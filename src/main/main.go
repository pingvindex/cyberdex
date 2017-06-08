package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

const webhookURL = "https://cyberdex.herokuapp.com/"

func main() {
	log.Println("Bot alives")
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI("253815575:AAHGADLTrRAx3P3sKFXGZ8Gd3Rh9o0IJgy8")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("Received text", update.Message.Text)

		switch update.Message.Text {
		case "привет":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "привет, не узнал тебя")
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "я пока очень туп. не знаю, что ответить ):")
		}

		bot.Send(message)
	}
}
