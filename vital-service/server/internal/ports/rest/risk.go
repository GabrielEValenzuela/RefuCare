package rest

import (
	"github.com/gofiber/fiber/v2"
)

// GetRiskByID godoc
// @Summary      Get risk score by patient ID
// @Description  Retrieve the risk status for a specific patient
// @Tags         risk
// @Produce      json
// @Param        id   path      string  true  "Patient ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /vitals/risk/{id} [get]
func GetRiskByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// TODO: Call external service or use analyzer logic to determine risk

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"patient_id": id,
		"risk":       "low", // static for now
	})
}
