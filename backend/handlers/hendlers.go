package handlers

import (
	"log"

	"github.com/come-on-readme/models"
	"github.com/come-on-readme/services"
	"github.com/come-on-readme/smash"
	"github.com/gofiber/fiber/v2"
)

//import "github.com/gofiber/fiber/v2"

func GenerateReadme(c *fiber.Ctx) error {
	var Contents []models.GeneralContent
	Contents, err := services.GetGeneralContents("yasirkelesh", "philo")
	if err != nil {
		log.Printf("error fetching repo info: %v", err)
		return err
	}
	smash.Smash(Contents)

	c.SendString(Contents[1].Name)
	return nil
}
