package analyze

import "flag"

func analyze() {
	// Get parameters
	logFilePath := flag.String("file", "", "Path to the log file")

	// print log

	// Initial channel
	var (
		logCh     = make(chan string, 1024)
		pvCh      = make(chan logData, 1024)
		uvCh      = make(chan logData, 1024)
		storageCh = make(chan storageBlock, 1024)
	)

	// log consumer

	// PV UV

	// store
}
