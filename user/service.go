package user

import (
    "golang.org/x/crypto/bcrypt"
    "company-project/employee"
)

type Service interface {
    CreateUser(userRequest UserRequest) (UserResponse, error)
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
