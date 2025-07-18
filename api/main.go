package main

import (
	"fmt"
	"log"
	"os"
	"url-shortener/middleware"
	"url-shortener/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/shorturls", routes.ShortenURL)
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
