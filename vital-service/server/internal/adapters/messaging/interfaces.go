package messaging

import "github.com/GabrielEValenzuela/RefuCare/internal/core/domain"

type VitalsPublisher interface {
	PublishVitals(*domain.Vitals) error
}
