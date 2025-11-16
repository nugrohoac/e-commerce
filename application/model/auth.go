package model

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
