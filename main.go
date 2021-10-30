package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"say-hi-backend/config"
	"say-hi-backend/controllers"
)
var dotEnvPath string

func init() {
	flag.StringVar(&dotEnvPath, "env", ".env", "Path to dotenv file")
	flag.Parse()
}

func main() {
	err := godotenv.Load(dotEnvPath)

	if err != nil {
		log.Fatal("Error while loading .env file. Please put near bin file")
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()

	app.Post("/messages", controllers.SaveMessage)
	app.Get("/messages/:id", controllers.GetOneMessageById)
	app.Listen(":" + os.Getenv("PORT"))
}
