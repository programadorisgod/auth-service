package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/programadorisgod/auth-service/src/models/user"
	authServices "github.com/programadorisgod/auth-service/src/services/auth"
)

func CreateUser(c *fiber.Ctx) error {
	var req user.UserRegister

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Email == "" || req.Pass == "" || req.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	u, error := authServices.FindUser(req.Email)

	if error != nil {
		return error
	}

	if u != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	id, err := authServices.SaveUser(&req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}
