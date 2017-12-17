package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

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
		case "привет":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "привет, не узнал тебя")
		case "/w":
			resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Moscow&APPID=7a3937709a28279ddeca2d281dec984f")
			if err != nil {
				log.Println("Данные о погоде недоступны")
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Данные о погоде недоступны")
			}
			message = tgbotapi.NewMessage(update.Message.Chat.ID, string(body))
		case "/dice":
			dice := strconv.Itoa(rand.Int()%6 + 1)
			message = tgbotapi.NewMessage(update.Message.Chat.ID, dice)
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "не знаю, что ответить ):")
		}

		bot.Send(message)
	}
}
