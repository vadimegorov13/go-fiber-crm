package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vadimegorov13/go-fiber-crm/pkg/common/config"
	"github.com/vadimegorov13/go-fiber-crm/pkg/common/db"
	"github.com/vadimegorov13/go-fiber-crm/pkg/lead"
)

func main() {
	dbHandler := db.InitDB()

	fiberConf := config.FiberConfig()
	logConf, file := config.LoggerConfig()
	defer file.Close()

	app := fiber.New(fiberConf)
	app.Use(logger.New(logConf))

	lead.RegisterRoutes(app, dbHandler)

	log.Fatal(app.Listen(":3000"))
}
