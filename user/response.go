package user

type UserResponse struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
    Phone int    `json:"phone"`
}