package web

import (
	"Cafeteria/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
)

type WebServices struct {
	Orchestrator *orchestrator.Orchestrator
}

func (ws *WebServices) Run() error {

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

}
