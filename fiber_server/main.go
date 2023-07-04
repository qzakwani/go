package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		log.Println(headers)
		return nil
	})
	app.Listen(":3000")
}
