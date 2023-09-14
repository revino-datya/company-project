package user

func ConvertToUserResponse(user User) UserResponse {
    return UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Name:  user.Employee.Name,
        Phone: user.Employee.Phone,
    }
}