package run

import (
	"fmt"
	"log"
	"os/exec"
)

// LL is command of  `ls -l [dir]`
func LL(dir string) {
	var cmd *exec.Cmd

	if dir == "" {
		cmd = exec.Command("ls", "-l")
	} else {
		cmd = exec.Command("ls", "-l", dir)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(out))
}
