package user

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
}

type LoginRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type UserUpdateRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Phone      int    `json:"phone"`
	Department uint   `json:"department"`
}
