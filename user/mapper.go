package user

func ConvertToUserResponse(u User) UserResponse {
	return UserResponse{
		Email:    u.Email,
		Password: u.Password,
	}
}
