package handlers

import (
	"log"

	"github.com/come-on-readme/services"
	"github.com/gofiber/fiber/v2"
)

//import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {

	message, err := services.FetchReadme("yasirkelesh", "MINISHELL")
	if err != nil {
		return err
	}

	err = services.GetSomething("yasirkelesh", "MINISHELL")
	if err != nil {
		log.Printf("error fetching repo info: %v", err)
		return err
	}

	//log.Printf("InfoMessage: %v", InfoMessage)

	/* 	something, err := services.GetSomething(c, "https://api.github.com/repos/yasirkelesh/MINISHELL/git/blobs/d1738830a94c34bc61576af7b08c52774f0672d7")
	   	log.Printf("something: %v", something)
	   	if err != nil {
	   		log.Printf("error: %v", err)
	   		return err
	   	} */
	log.Printf("message: %v", message)
	//c.SendString(message)
	return nil
}
