package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"

	"github.com/programadorisgod/auth-service/src/config"
	"github.com/programadorisgod/auth-service/src/models/user"
)

func Register(c *fiber.Ctx) error {

	var req user.UserRegister

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	postBody, err := json.Marshal(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error, try later",
		})
	}

	urlUsers, err := url.JoinPath(config.Url_users, "register")
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
			"error": "Error reading response from user service",
		})
	}
	//TODO: generate token and send token
	return c.Status(res.StatusCode).Send(body)
}
