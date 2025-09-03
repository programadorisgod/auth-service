package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/programadorisgod/auth-service/src/config"
	"github.com/programadorisgod/auth-service/src/controllers/auth"
)

func main() {

	config.InitDB()
	defer config.DB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm healthy")
	})

	app.Post("/login", auth.Login)
	app.Post("/register", auth.CreateUser)

	log.Fatal(app.Listen(":4000"))

}
