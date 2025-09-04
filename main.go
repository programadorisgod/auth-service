package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/programadorisgod/auth-service/src/config"
	"github.com/programadorisgod/auth-service/src/controllers/auth"
)

func main() {

	config.InitDB()
	defer config.DB.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm healthy")
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Post("/login", auth.Login)
	app.Post("/register", auth.CreateUser)

	log.Fatal(app.Listen(":4000"))

}
