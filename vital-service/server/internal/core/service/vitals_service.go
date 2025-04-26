package service

import (
	"fmt"
	"time"

	"github.com/GabrielEValenzuela/RefuCare/internal/adapters/messaging"
	grpcport "github.com/GabrielEValenzuela/RefuCare/internal/ports/grpc"

	"github.com/GabrielEValenzuela/RefuCare/internal/core/domain"
	"github.com/google/uuid"
)

type VitalsService struct {
	repo       VitalsRepository
	grpcClient *grpcport.PatientRecordsClient
	publisher  messaging.VitalsPublisher
}

func NewVitalsService(repo VitalsRepository, grpc *grpcport.PatientRecordsClient, pub messaging.VitalsPublisher) *VitalsService {
	return &VitalsService{
		repo:       repo,
		grpcClient: grpc,
		publisher:  pub,
	}
}

func (s *VitalsService) RecordVitals(input domain.Vitals) (*domain.Vitals, error) {
	input.ID = uuid.New().String()
	input.RecordedAt = time.Now().UTC()

	if input.Systolic <= 0 || input.Diastolic <= 0 || input.PatientID == "" {
		return nil, ErrInvalidVitals
	}

	err := s.repo.SaveVitals(input)
	if err != nil {
		return nil, err
	}

	// Future: publish to MQ, trigger risk analysis

	return &input, nil
}

func (s *VitalsService) CalculateRiskByID(id string) (map[string]interface{}, error) {
	v, err := s.repo.GetVitalsByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not find vitals: %w", err)
	}

	pinfo, err := s.grpcClient.GetPatientInfo(v.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch patient info: %w", err)
	}

	// Simple rule-based risk calculation
	risk := "low"
	if v.Systolic > 140 || v.Diastolic > 90 || v.HeartRate > 100 || v.Temp > 38.0 {
		risk = "medium"
	}
	if pinfo.HasDiabetes && v.Systolic > 160 {
		risk = "high"
	}

	return map[string]interface{}{
		"patient_id":   v.PatientID,
		"name":         pinfo.Name,
		"age":          pinfo.Age,
		"has_diabetes": pinfo.HasDiabetes,
		"systolic":     v.Systolic,
		"diastolic":    v.Diastolic,
		"heart_rate":   v.HeartRate,
		"temperature":  v.Temp,
		"risk":         risk,
		"recorded_at":  v.RecordedAt,
	}, nil
}

func (s *VitalsService) FindVitalsByID(id string) (*domain.Vitals, error) {
	return s.repo.GetVitalsByID(id)
}

// Define app-level error
var ErrInvalidVitals = &ServiceError{"invalid vitals input"}

type ServiceError struct {
	Msg string
}

func (e *ServiceError) Error() string {
	return e.Msg
}
