package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/programadorisgod/auth-service/src/config"
	"github.com/programadorisgod/auth-service/src/models/user"
)

func Login(c *fiber.Ctx) error {
	var req user.UserLogin

	var userWrapper struct {
		User user.User `json:"user"`
	}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if req.Email == "" || req.Pass == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	postBody, err := json.Marshal(req)

	log.Print("post body", string(postBody))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error, try later",
		})
	}

	urlUsers, err := url.JoinPath(config.Url_users, "login")

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to build user service URL",
		})
	}

	res, err := http.Post(urlUsers, "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error reading  AI user service",
		})
	}

	if res.StatusCode == http.StatusUnauthorized {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	if res.StatusCode == http.StatusNotFound {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
	}

	if res.StatusCode != http.StatusOK {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "unexpected error from user service"})
	}
	log.Print(string(body))

	if err := json.Unmarshal(body, &userWrapper); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "invalid response from user service"})
	}

	log.Print(userWrapper.User)

	return c.Status(res.StatusCode).JSON(fiber.Map{"user": userWrapper.User})
}
