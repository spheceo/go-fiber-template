package handler

import (
	"net/http"
	
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func favicon(c fiber.Ctx) error {
	return c.SendFile("./public/favicon.ico")
}

func index(c fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "Welcome to {{ project_name }}!",
	})
}

// HTTP Handler which Vercel looks for
func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	// CORS Setup
	app.Use(cors.New())

	// Define Routes
	app.Get("/", index)
	app.Get("/favicon.ico", favicon)

	// Serve HTTP
	http.HandlerFunc(adaptor.FiberApp(app)).ServeHTTP(w, r)
}
