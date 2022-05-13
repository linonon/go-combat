package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"os"

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
		if update.Message != nil {
			// 看進來的是什麼指令
			cmd := update.Message.Command()
			stdout := ""
			switch cmd {
			case "shell":
				fmt.Println("exec shell...")
				shellcmd := update.Message.CommandArguments()
				bannedCmds := []string{"rm", "mv", "cp"}
				for _, cmd := range bannedCmds {
					if strings.Contains(shellcmd, cmd) {
						bot.Send(
							tgb.NewMessage(update.Message.Chat.ID, fmt.Sprintf("'%s' command is banned", cmd)),
						)
					}
				}
				fmt.Printf("shellcmd: '%s'\n", shellcmd)
				stdoutBytes, err := exec.Command("bash", "-c", shellcmd).CombinedOutput()
				if err != nil {
					fmt.Println(err)
				}
				stdout = string(stdoutBytes)
				fmt.Println("stdout: ", stdout)
			case "md":
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				m := update.Message.CommandArguments()

				msg := tgb.NewMessage(update.Message.Chat.ID, "**"+m+"**")
				// msg.ReplyToMessageID = update.Message.MessageID
				msg.Entities = append(msg.Entities, tgb.MessageEntity{Type: "bold"})

				bot.Send(msg)

			default:
				fmt.Println("Do nothing")
			}
			msg := tgb.NewMessage(update.Message.Chat.ID, stdout)

			bot.Send(msg)
		}
	}
}

func main() {
	easyReply()
}
