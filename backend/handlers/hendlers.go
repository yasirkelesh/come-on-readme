package handlers

import (
	"github.com/come-on-readme/services"
	"github.com/gofiber/fiber/v2"
)

//import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	message, err := services.FetchReadme("yasirkelesh", "INCEPTION")
	if err != nil {
		return err
	}

	c.SendString(message)
	return nil
}
