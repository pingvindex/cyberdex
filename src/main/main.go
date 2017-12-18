package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"weather"

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

		incoming := strings.Split(update.Message.Text, " ")

		switch incoming[0] {
		case "/w", "/weather":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, weather.GetWeather())
		case "/d", "/dice":
			dice := strconv.Itoa(rand.Int()%6 + 1)
			message = tgbotapi.NewMessage(update.Message.Chat.ID, dice)
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID,
				"Команды для бота:\n\t/w, /weather\t\tПоказать погоду в Москве\n\t/d, /dice\t\tБросить кость, результат от 1 до 6")
		}

		bot.Send(message)
	}
}
