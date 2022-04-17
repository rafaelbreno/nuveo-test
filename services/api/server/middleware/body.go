package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CheckBody if was received a body,
// when Unmarshalling form a empty
// byte array, it panics.
func CheckBody(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": "empty body",
			})
	}

	return c.Next()
}
