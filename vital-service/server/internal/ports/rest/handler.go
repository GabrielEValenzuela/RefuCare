package rest

import (
	"github.com/GabrielEValenzuela/RefuCare/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, vitalsSvc *service.VitalsService) {
	v := app.Group("/vitals")

	// REST controller logic can now call business logic via vitalsSvc
	v.Post("/", CreateVitalsHandler(vitalsSvc))
	v.Get("/:id", GetVitalsByIDHandler(vitalsSvc))
	v.Get("/risk/:id", GetRiskByIDHandler(vitalsSvc))
}
