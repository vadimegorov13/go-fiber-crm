package config

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {

	fiberConf := fiber.Config{
		AppName:     "Simple CRM using fiber",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	return fiberConf
}
