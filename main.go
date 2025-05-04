package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/suraboy/poc-go-scalar/docs"
	"github.com/yokeTH/go-pkg/scalar"
)

// @title My API
// @version 1.0
// @description My test API for Fiber + Scalar
func main() {
	app := fiber.New()

	// setup scalar
	app.Use(scalar.New(scalar.Config{
		BasePath: "/",
		Path:     "swagger",
		Title:    "POC Scalar Documentation",
	}))

	app.Get("/hello", HelloHandler)

	app.Listen(":3000")
}

// HelloHandler godoc
// @Summary Say Hello
// @Description Responds with Hello World
// @Tags example
// @Success 200 {string} string "ok"
// @Router /hello [get]
func HelloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
