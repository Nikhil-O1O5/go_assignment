package models

type UserRequest struct {
	Name string `json:"name" validate:"required"` 
	DOB  string `json:"dob" validate:"required"`  
}

type UserResponse struct {
	ID   int32  `json:"id"`   // [cite: 39]
	Name string `json:"name"` // [cite: 41]
	DOB  string `json:"dob"`  // [cite: 42]
	Age  int    `json:"age"`  // Calculated dynamically [cite: 50]
}