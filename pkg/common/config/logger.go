package config

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() logger.Config {
	file, err := os.OpenFile("../logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	logConf := logger.Config{
		Output: file,
	}

	return logConf
}
