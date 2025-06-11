package dto

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupDTO struct {
	Phone string `json:"phone"`
	LoginDTO
}

type VerificationCodeDTO struct {
	Code int `json:"code"`
}
