package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	Gender   uint8  `json:"gender"`
	Age      uint8  `json:"age"`
	Email    string `json:"email" binding:"required,min=8,max=30,email"`
}

type UpdateUserRequest struct {
	Id       uuid.UUID `json:"id" binding:"required,uuid"`
	Password *string   `json:"password"`
	Gender   *uint8    `json:"gender"`
	Age      *uint8    `json:"age"`
	Email    *string   `json:"email" binding:"omitempty,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
