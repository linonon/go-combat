package main

import (
	"fmt"
	"os"
	"timed-task/pkg/run"
)

func main() {

	run.LL("")
	os.Chdir("~/Workspace")
	dir, _ := os.Getwd()
	os.Chdir("~/Workspace")
	fmt.Println(dir)
	run.LL("")
}
