package command

import (
	"fmt"
	"log"

	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// MD will return the message in markdown format.
func MD(update tgb.Update) tgb.MessageConfig {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	fmt.Println(update.Message.From.LanguageCode)

	m := update.Message.CommandArguments()

	msg := tgb.NewMessage(update.Message.Chat.ID, "")
	// msg.ReplyToMessageID = update.Message.MessageID
	msg.Text = m
	msg.Entities = append(msg.Entities, tgb.MessageEntity{Type: "bold", Offset: 0, Length: len(m), Language: "zh-hans"})

	fmt.Println(msg.Text, msg.Entities)
	return msg
}
