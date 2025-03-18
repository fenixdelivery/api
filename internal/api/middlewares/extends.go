package middlewares

import (
	fiber "github.com/gofiber/fiber/v2"
)

func ExtendsContext[T any](dependency T, key string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(key, dependency)
		return c.Next()
	}
}