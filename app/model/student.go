package model

type Student struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Dob   string `json:"dob,omitempty"`
	Gpa   float64 `json:"gpa,omitempty"`
}
