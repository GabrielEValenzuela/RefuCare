package rest

import (
	"github.com/GabrielEValenzuela/RefuCare/internal/core/domain"
	"github.com/GabrielEValenzuela/RefuCare/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

type CreateVitalsRequest struct {
	PatientID string  `json:"patient_id" validate:"required"`
	Systolic  float64 `json:"systolic" validate:"required"`
	Diastolic float64 `json:"diastolic" validate:"required"`
	HeartRate int     `json:"heart_rate"`
	Temp      float64 `json:"temperature"`
}

// CreateVitals godoc
// @Summary      Submit new vitals
// @Description  Creates a new vitals record for a patient
// @Tags         vitals
// @Accept       json
// @Produce      json
// @Param        vitals  body  CreateVitalsRequest  true  "Vitals Input"
// @Success      201  {object}  domain.Vitals
// @Failure      400  {object}  map[string]string
// @Router       /vitals [post]
func CreateVitalsHandler(svc *service.VitalsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateVitalsRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}

		v := domain.Vitals{
			PatientID: req.PatientID,
			Systolic:  req.Systolic,
			Diastolic: req.Diastolic,
			HeartRate: req.HeartRate,
			Temp:      req.Temp,
		}

		saved, err := svc.RecordVitals(v)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(saved)
	}
}

// GetVitalsByIDHandler godoc
// @Summary      Get vitals by ID
// @Description  Fetch a single vitals record by UUID
// @Tags         vitals
// @Produce      json
// @Param        id   path      string  true  "Vitals ID"
// @Success      200  {object}  domain.Vitals
// @Failure      404  {object}  map[string]string
// @Router       /vitals/{id} [get]
func GetVitalsByIDHandler(svc *service.VitalsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		v, err := svc.FindVitalsByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "vitals not found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(v)
	}
}

// GetRiskByIDHandler godoc
// @Summary      Get risk assessment by vitals ID
// @Description  Calculate risk based on vitals and patient info
// @Tags         vitals
// @Produce      json
// @Param        id   path      string  true  "Vitals ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]string
// @Router       /vitals/risk/{id} [get]
func GetRiskByIDHandler(svc *service.VitalsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := svc.CalculateRiskByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(result)
	}
}
