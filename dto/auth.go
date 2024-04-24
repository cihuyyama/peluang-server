package dto

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type AuthResponse struct {
}
