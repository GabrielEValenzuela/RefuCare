package domain

type Patient struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"` // format: YYYY-MM-DD
	Gender    string `json:"gender"`
}
