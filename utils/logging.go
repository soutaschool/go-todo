package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ファイルの追加、作成、読み込みのどれかを行なう
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
