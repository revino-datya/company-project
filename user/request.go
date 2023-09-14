package user

type UserRequest struct {
    ID       uint   `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
    Phone    int    `json:"phone"`
}
