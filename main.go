package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	apiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if apiToken == "" {
		log.Fatalln("Environment variable TELEGRAM_BOT_API_TOKEN must be defined")
	}
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Fatalln(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updatesChan, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting handler loop")
	for update := range updatesChan {
		log.Printf("received message from: %v, with text: %v\n", update.Message.From.UserName, update.Message.Text)
		if update.Message.IsCommand() {
			log.Println("message is a command")
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand you")
			_, err := bot.Send(reply)
			if err != nil {
				log.Println(err)
			}
			continue
		}
		reply := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}
}
