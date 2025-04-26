package service

import "github.com/GabrielEValenzuela/RefuCare/internal/core/domain"

type VitalsRepository interface {
	SaveVitals(v domain.Vitals) error
	GetVitalsByID(id string) (*domain.Vitals, error)
}
