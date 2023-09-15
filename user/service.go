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
	FindAllUsers() ([]UserResponse, error)
	FindUserByID(ID uint) (UserResponse, error)
	UpdateUser(ID uint, userUpdateRequest UserUpdateRequest) (UserResponse, error)
	DeleteUser(ID uint) error
	Login(loginRequest LoginRequest) (string, error)
}

type service struct {
	repository Repository
	// employeeRepo employee.Repository
}

// func NewService(repository Repository, employeeRepo employee.Repository) Service {
func NewService(repository Repository) Service {
	// return &service{repository, employeeRepo}
	return &service{repository}
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

// func (s *service) UpdateUser(userID uint, userRequest UserRequest) (UserResponse, error) {
// 	// Cari pengguna berdasarkan ID
// 	existingUser, err := s.repository.FindByID(userID)
// 	if err != nil {
// 		return UserResponse{}, err
// 	}

// 	// Jika userRequest.Password tidak kosong, maka kita perlu menghash password yang baru
// 	if userRequest.Password != "" {
// 		// Generate hashed password
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			return UserResponse{}, err
// 		}
// 		// Update password
// 		existingUser.Password = string(hashedPassword)
// 	}

// 	// Update data lainnya
// 	existingUser.Email = userRequest.Email
// 	existingUser.Employee.Name = userRequest.Name
// 	existingUser.Employee.Phone = userRequest.Phone
// 	// existingUser.Employee.DepartmentID = userRequest.Department

// 	// Simpan perubahan ke database
// 	updatedUser, err := s.repository.Update(existingUser)
// 	if err != nil {
// 		return UserResponse{}, err
// 	}

// 	// Gunakan mapper untuk mengonversi User entity yang diperbarui ke UserResponse
// 	userResponse := ConvertToUserResponse(updatedUser)

// 	return userResponse, nil
// }

func (s *service) UpdateUser(userID uint, userUpdateRequest UserUpdateRequest) (UserResponse, error) {
	// Cari pengguna berdasarkan ID
	existingUser, err := s.repository.FindByID(userID)
	if err != nil {
		return UserResponse{}, err
	}

	// Update data pengguna (Email, Password jika ada)
	existingUser.Email = userUpdateRequest.Email
	if userUpdateRequest.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userUpdateRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			return UserResponse{}, err
		}
		existingUser.Password = string(hashedPassword)
	}

	// Update data Employee
	existingUser.Employee.Name = userUpdateRequest.Name
	existingUser.Employee.Phone = userUpdateRequest.Phone
	// Jika Anda ingin mengupdate DepartmentID juga, Anda dapat melakukannya di sini
	existingUser.Employee.DepartmentID = &userUpdateRequest.Department

	// Simpan perubahan ke database
	updatedUser, err := s.repository.Update(existingUser)
	if err != nil {
		return UserResponse{}, err
	}

	// Gunakan mapper untuk mengonversi User entity yang diperbarui ke UserResponse
	userResponse := ConvertToUserResponse(updatedUser)

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

// func (s *service) DeleteUser(userID uint) error {
// 	// Cari pengguna berdasarkan ID
// 	userToDelete, err := s.repository.FindByID(userID)
// 	if err != nil {
// 		return err // Return error if user is not found
// 	}

// 	// Hapus pengguna dari database
// 	err = s.repository.Delete(userToDelete)
// 	if err != nil {
// 		return err // Return error if there was an issue deleting the user
// 	}

// 	return nil // Return nil if user is successfully deleted
// }

func (s *service) DeleteUser(userID uint) error {
	// Cari pengguna berdasarkan ID
	existingUser, err := s.repository.FindByID(userID)
	if err != nil {
		return err
	}

	// Hapus pengguna dari database
	err = s.repository.Delete(existingUser)
	if err != nil {
		return err
	}

	// Jika pengguna telah berhasil dihapus, Anda juga dapat menghapus data Employee yang sesuai
	// Ini hanya akan berfungsi jika ada hubungan referensial antara User dan Employee,
	// dan database Anda mendukung cascading delete.
	// Jika tidak, Anda harus menghapus Employee terpisah setelah menghapus User.

	return nil
}
