package domain

import "time"

type Vitals struct {
	ID         string    `json:"id"`          // UUID or hash
	PatientID  string    `json:"patient_id"`  // foreign key to patient
	Systolic   float64   `json:"systolic"`    // mmHg
	Diastolic  float64   `json:"diastolic"`   // mmHg
	HeartRate  int       `json:"heart_rate"`  // bpm
	Temp       float64   `json:"temperature"` // Celsius
	RecordedAt time.Time `json:"recorded_at"` // UTC
}
