package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		log.Printf("Incoming Request: %s %s", c.Method(), c.Path())

		err := c.Next()

		duration := time.Since(start)
		log.Printf("Request %s %s took %v and responded with status %d",
			c.Method(),
			c.Path(),
			duration,
			c.Response().StatusCode(),
		)

		return err
	}
}
