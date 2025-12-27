package request

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   *uint8 `json:"gender"`
	Age      *uint8 `json:"age"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	Id       string  `json:"id"`
	Password *string `json:"password"`
	Gender   *uint8  `json:"gender"`
	Age      *uint8  `json:"age"`
	Email    *string `json:"email"`
}
