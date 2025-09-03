package auth

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/programadorisgod/auth-service/src/models/user"
	authServices "github.com/programadorisgod/auth-service/src/services/auth"
)

func Login(c *fiber.Ctx) error {
	var req user.UserLogin

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	log.Println(req.Email)
	log.Println(req.Pass)

	if req.Email == "" || req.Pass == " " {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	u, err := authServices.FindUser(req.Email)

	if err != nil {
		return err
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	if u.Email == req.Email && u.Pass == req.Pass {
		return c.JSON(fiber.Map{
			"user": u,
		})
	}

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"error": "invalid credentials",
	})
}
