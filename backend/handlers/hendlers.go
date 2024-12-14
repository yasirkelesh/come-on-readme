package handlers

import (
	"log"
	"strings"

	"github.com/come-on-readme/models"
	"github.com/come-on-readme/services"
	"github.com/gofiber/fiber/v2"
)

//import "github.com/gofiber/fiber/v2"

func GenerateReadme(c *fiber.Ctx) error {
	var Contents []models.GeneralContent
	Contents, err := services.GetContents("yasirkelesh", "MINISHELL")
	if err != nil {
		log.Printf("error fetching repo info: %v", err)
		return err
	}

	for _, content := range Contents {
		switch content.Type {
		case "file":
			if strings.Contains(strings.ToLower(content.Name), "main") || 
			strings.Contains(strings.ToLower(content.Name), "readme") || 
			strings.Contains(strings.ToLower(content.Name), "makefile") ||
			strings.Contains(strings.ToLower(content.Name), ".yml")||
			strings.Contains(strings.ToLower(content.Name), ".sh") {
				log.Printf("File: %s", content.Name)
			}
		case "dir":
			log.Printf("Dir: %s", content.Name)
		}
	}

	c.SendString(Contents[1].Name)
	return nil
}
