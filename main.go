package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./index.html")
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
