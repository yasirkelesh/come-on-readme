package smash

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/come-on-readme/models"
	"github.com/come-on-readme/services"
)

func SmashFile(name, content, encoding string) (string, error) {
	switch encoding {
	case "base64":
		data := base64Decode(content)
		return data, nil
	default:
		return "", fmt.Errorf("encoding not supported: %s", encoding)
	}
	//return content, nil
}

func Smash(Contents *[]models.GeneralContent) {
	for _, content := range *Contents {
		switch content.Type {
		case "file":
			if strings.Contains(strings.ToLower(content.Name), "main") ||
				strings.Contains(strings.ToLower(content.Name), "readme") ||
				strings.Contains(strings.ToLower(content.Name), "makefile") ||
				strings.Contains(strings.ToLower(content.Name), ".yml") ||
				strings.Contains(strings.ToLower(content.Name), ".sh") ||
				strings.Contains(strings.ToLower(content.Name), "docker") {

				var FileContent *models.FileContent
				FileContent, err := services.GetFile(content.URL)
				if err != nil {
					log.Printf("error fetching repo info: %v", err)
					return
				}

				data, err := SmashFile(content.Name, FileContent.Content, FileContent.Encoding)
				if err != nil {
					log.Printf("error fetching repo info: %v", err)
					return
				}
				fmt.Println(data)

			}
		case "dir":
			/* if strings.Contains(strings.ToLower(content.Name), "src") ||
			strings.Contains(strings.ToLower(content.Name), "include") ||
			strings.Contains(strings.ToLower(content.Name), "lib") ||
			strings.Contains(strings.ToLower(content.Name), "obj") ||
			strings.Contains(strings.ToLower(content.Name), "hendlers") ||
			strings.Contains(strings.ToLower(content.Name), "main")  */{

				var Contents *[]models.GeneralContent
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

func base64Decode(data string) string {
	// Split the data by newline
	lines := strings.Split(data, "\n")

	// List to hold Base64 encoded lines
	var decodedLines []string
	for _, line := range lines {
		decoded, err := base64.StdEncoding.DecodeString(line)
		if err != nil {
			log.Printf("error decoding base64: %v", err)
			return ""
		}
		decodedLines = append(decodedLines, string(decoded))
	}

	// Combine the encoded lines into a single string
	result := strings.Join(decodedLines, "")

	return result
}
