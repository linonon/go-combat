package main

import (
	"fmt"
	"log"

	"os"

	"tgbot/command"

	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func easyReply() {
	bot, err := tgb.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgb.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			var msg tgb.MessageConfig

			switch update.Message.Command() {
			case "shell":
				msg = command.Shell(update)
			case "md":
				msg = command.MD(update)
			default:
				fmt.Println("Do nothing")
			}

			bot.Send(msg)
		}
	}
}

func main() {
	easyReply()
}
