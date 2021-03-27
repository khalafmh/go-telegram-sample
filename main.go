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
		if update.Message == nil {
			log.Printf("Received empty update with id %v\n", update.UpdateID)
			continue
		}
		var reply tgbotapi.MessageConfig
		if update.Message.IsCommand() {
			log.Printf(
				"Received message from: %v, with command: %v, and arguments: %v\n",
				update.Message.From.UserName,
				update.Message.Command(),
				update.Message.CommandArguments(),
			)
			reply = tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand you")
		} else {
			log.Printf(
				"Received message from: %v, with text: %v\n",
				update.Message.From.UserName,
				update.Message.Text,
			)
			reply = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		}
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}
}
