package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() (logger.Config, *os.File) {
	logsDir := "logs"
	os.Mkdir(logsDir, os.ModePerm) // create the logs directory if it doesn't exist

	now := time.Now()
	filename := fmt.Sprintf("%s/%d-%02d-%02d.log", logsDir, now.Year(), now.Month(), now.Day())
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logConf := logger.Config{
		Output: file,
	}

	return logConf, file
}
