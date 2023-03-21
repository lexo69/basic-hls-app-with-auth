package app

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Host variable. Used for generating video url and as a fiber listening insterface
var HOST = "127.0.0.1:9000"

// Variable for a global fiber instance
var HTTP *fiber.App

// Set up fiber
func InitHTTP() {

	HTTP = fiber.New(fiber.Config{
		Prefork:               false,
		ServerHeader:          "hls-streamer",
		CaseSensitive:         true,
		DisableStartupMessage: true,
		ErrorHandler:          AppErrorHandler,
	})

	//Test route
	HTTP.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": "server is up and running",
		})
	})

}

func AppErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	err = c.Status(code).JSON(fiber.Map{
		"error": err,
	})
	if err != nil {
		// In case the SendFile fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	// Return from handler
	return nil
}
