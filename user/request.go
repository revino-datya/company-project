package user

type UserRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
    Phone    int    `json:"phone"`
}
