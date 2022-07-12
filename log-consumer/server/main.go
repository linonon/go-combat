package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type cmdParams struct {
	logFilePath string
	routineNum  int
}
type digData struct {
	timestamp string
	dType     string
}
type logData struct {
	data digData
	uid  string
}
type storageBlock struct {
	counterType  string
	storageModel string
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

}

func main() {
	// Get parameters
	logFilePath := flag.String("file", "./track.json", "Path to the log file")
	routineNum := flag.Int("routineNum", 5, "consumerNumber")
	l := flag.String("l", "./log.txt", "storage log")

	params := cmdParams{
		logFilePath: *logFilePath,
		routineNum:  *routineNum,
	}
	// print log
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Out = logFd
	defer logFd.Close()

	log.Infoln("Exec start.")
	log.Infof("Params: logFilePath=%s, routineNum=%d", params.logFilePath, params.routineNum)

	// Initial channel
	var (
		logCh     = make(chan string, 3*params.routineNum)
		pvCh      = make(chan logData, params.routineNum)
		uvCh      = make(chan logData, params.routineNum)
		storageCh = make(chan storageBlock, params.routineNum)
	)

	// log consumer
	go readFile(&params, logCh)

	// mutiple process log
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logCh, pvCh, uvCh)
	}

	// PV UV
	go pvCounter(pvCh, storageCh)
	go uvCounter(uvCh, storageCh)
	// go xxxCounter(xxxch, storageCh)
	// go xxxCounter(xxxch, storageCh)
	// go xxxCounter(xxxch, storageCh)

	// store
	go dataStorage(storageCh)

	time.Sleep(1000 * time.Second)
}

func readFile(params *cmdParams, logCh chan string) {
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("ReadFile can't open file: %v", err)
	}
	defer fd.Close()

	count := 0
	bufferRead := bufio.NewReader(fd)
	for {
		line, err := bufferRead.ReadString('\n')
		logCh <- line
		a := <-logCh
		count++
		if count%(1000*params.routineNum) == 0 {
			log.Infof("ReadLine to %v", count)
			log.Infof("ReadLine is %v", a)
		}
		if err != nil {
			if err == io.EOF {
				time.Sleep(3 * time.Second)
				log.Infof("ReadLine waiting...")
			} else {
				log.Warningf("ReadFile error: %v", err)
			}
		}
	}

}

func logConsumer(logCh chan string, pvCh chan logData, uvCh chan logData) {
	// for logStr := range logCh {
	// 	// 課: 切割日至字符串, 扣出打點上報的數據
	// 	// 我: 轉化成 Json 格式
	// 	pvCh <- logStr
	// 	uvCh <- logStr
	// }
}

func pvCounter(pvCh chan logData, storageCh chan storageBlock) {

}

func uvCounter(uvCh chan logData, storageCh chan storageBlock) {}

func dataStorage(storageCh chan storageBlock) {}
