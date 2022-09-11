package web

import (
	"Cafeteria/pkg/common/services/io/web/controllers"
	"Cafeteria/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type WebServices struct {
	Orchestrator *orchestrator.Orchestrator
}

func (ws *WebServices) Run() error {

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	(&controllers.UserController{UseCases: ws.Orchestrator}).Init(v1)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// Restricted Routes
	app.Get("/restricted", restricted)

	//Hello
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	/////////Here must be controllers

	/////////END Here must be controllers

	return app.Listen(":3000")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
