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
	config.LoadaEnv()
	config.InitDB()
	defer config.DB.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: false,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm healthy!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Post("/api/auth/register", auth.Register)
	app.Post("/api/auth/login", auth.Login)

	log.Fatal(app.Listen("0.0.0.0:4000"))

}
