package command

import (
	"fmt"
	"os/exec"
	"strings"

	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Shell will run the given command
func Shell(update tgb.Update) tgb.MessageConfig {
	argus := update.Message.CommandArguments()
	bannedCmds := []string{"rm", "mv", "cp"}

	for _, cmd := range bannedCmds {
		if strings.Contains(argus, cmd) {
			return tgb.NewMessage(update.Message.Chat.ID, fmt.Sprintf("'%s' command is banned", cmd))
		}
	}

	fmt.Printf("shellcmd: '%s'\n", argus)

	stdoutBytes, err := exec.Command("bash", "-c", argus).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	stdout := string(stdoutBytes)
	fmt.Println("stdout: ", stdout)
	return tgb.NewMessage(update.Message.Chat.ID, stdout)
}
