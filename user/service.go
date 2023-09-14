package user

import (
	"company-project/employee"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"os"
	"time"
)

type Service interface {
	CreateUser(userRequest UserRequest) (UserResponse, error)
	Login(loginRequest LoginRequest) (string, error)
}

type service struct {
	repository   Repository
	employeeRepo employee.Repository
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

func (s *service) Login(loginRequest LoginRequest) (string, error) {
	//get user
	user, err := s.repository.FindByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	} else if user.ID == 0 {
		return "", errors.New("Invalid email or password")
	}
	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(loginRequest.Password))
	if err != nil {
		return "", err
	}
	//sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
