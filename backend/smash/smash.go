package smash

import (
	"log"
	"strings"

	"github.com/come-on-readme/models"
	"github.com/come-on-readme/services"
)

func SmashFileAndBlob(name, content, encoding string) string {
	return name
}

func Smash(Contents []models.GeneralContent) {
	for _, content := range Contents {
		switch content.Type {
		case "file":
			if strings.Contains(strings.ToLower(content.Name), "main") ||
				strings.Contains(strings.ToLower(content.Name), "readme") ||
				strings.Contains(strings.ToLower(content.Name), "makefile") ||
				strings.Contains(strings.ToLower(content.Name), ".yml") ||
				strings.Contains(strings.ToLower(content.Name), ".sh") {

				name := SmashFileAndBlob(content.Name, content.URL, "")
				log.Printf("File: %s", name)
			}
		case "dir":
			/* if strings.Contains(strings.ToLower(content.Name), "src") ||
			strings.Contains(strings.ToLower(content.Name), "include") ||
			strings.Contains(strings.ToLower(content.Name), "lib") ||
			strings.Contains(strings.ToLower(content.Name), "obj") ||
			strings.Contains(strings.ToLower(content.Name), "hendlers") ||
			strings.Contains(strings.ToLower(content.Name), "main")  */{

				var Contents []models.GeneralContent
				Contents, err := services.GetDir(content.URL)
				if err != nil {
					log.Printf("error fetching repo info: %v", err)
					return
				}
				Smash(Contents)

			}
		}
	}
}
