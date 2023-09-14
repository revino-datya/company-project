package user

import (
    "golang.org/x/crypto/bcrypt"
    "company-project/employee"
)

type Service interface {
    CreateUser(userRequest UserRequest) (UserResponse, error)
    FindAllUsers() ([]UserResponse, error)
    FindUserByID(ID uint) (UserResponse, error)
    // UpdateUser(ID uint, userRequest UserRequest) (UserResponse, error)
    // DeleteUser(ID uint) error
}

type service struct {
    repository     Repository
    employeeRepo   employee.Repository
}

func NewService(repository Repository, employeeRepo employee.Repository) Service {
    return &service{repository, employeeRepo}
}

func (s *service) CreateUser(userRequest UserRequest) (UserResponse, error) {
    // Generate hashed password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
    if err != nil {
        return UserResponse{}, err
    }

    // Create User entity
    newUser := User{
        Email:    userRequest.Email,
        Password: string(hashedPassword),
    }

    // Create Employee entity with null DepartmentID
    newEmployee := employee.Employee{
        Name:         userRequest.Name,
        Phone:        userRequest.Phone,
        DepartmentID: nil,
    }

    // Associate Employee with User
    newUser.Employee = newEmployee

    // Save User to the database
    createdUser, err := s.repository.Create(newUser)
    if err != nil {
        return UserResponse{}, err
    }

    // Use a mapper to transform the User entity to UserResponse
    userResponse := ConvertToUserResponse(createdUser)

    return userResponse, nil
}

func (s *service) FindAllUsers() ([]UserResponse, error) {
    // Dapatkan semua pengguna dari penyimpanan (misalnya database)
    users, err := s.repository.FindAll()
    if err != nil {
        return nil, err
    }

    // Buat sebuah slice untuk menyimpan UserResponse
    userResponses := make([]UserResponse, len(users))

    // Map setiap entitas User ke UserResponse
    for i, user := range users {
        userResponses[i] = ConvertToUserResponse(user)
    }

    return userResponses, nil
}

func (s *service) FindUserByID(userID uint) (UserResponse, error) {
    // Dapatkan pengguna berdasarkan ID dari penyimpanan (misalnya database)
    user, err := s.repository.FindByID(userID)
    if err != nil {
        return UserResponse{}, err
    }

    // Konversi entitas User ke UserResponse menggunakan mapper
    userResponse := ConvertToUserResponse(user)

    return userResponse, nil
}


// func (s *service) UpdateUser(updateRequest UserRequest) (UserResponse, error) {
//     // Dapatkan pengguna berdasarkan ID dari penyimpanan (misalnya database)
//     user, err := s.repository.FindByID(updateRequest.ID)
//     if err != nil {
//         return UserResponse{}, err
//     }

//     // Update data pengguna dengan nilai yang diterima dari updateRequest
//     user.Email = updateRequest.Email
//     user.Name = updateRequest.Name
//     user.Phone = updateRequest.Phone

//     // Simpan perubahan ke penyimpanan
//     if err := s.repository.Update(user); err != nil {
//         return UserResponse{}, err
//     }

//     // Konversi entitas User yang telah diupdate ke UserResponse menggunakan mapper
//     userResponse := ConvertToUserResponse(user)

//     return userResponse, nil
// }

// func (s *service) DeleteUser(ID uint) error {
//     // Implement the logic to delete a user by ID
// }
