package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"company-project/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var userRequest user.UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})

			return

		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	// Create user using the service
	userResponse, err := h.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	// Mendapatkan ID pengguna dari URL atau permintaan
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid user ID",
		})
		return
	}

	// Membaca data permintaan JSON
	var userRequest user.UserUpdateRequest
	err = c.ShouldBindJSON(&userRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return

		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	// Menggunakan service untuk melakukan pembaruan pengguna
	userResponse, err := h.userService.UpdateUser(uint(userID), userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	userResponses, err := h.userService.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponses,
	})
}

// func (h *userHandler) GetUserById(c *gin.Context) {
// 	userIDParam := c.Param("id")

// 	// Ubah ID pengguna menjadi tipe data yang sesuai (misalnya uint)
// 	userID, err := strconv.ParseUint(userIDParam, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": "Invalid user ID",
// 		})
// 		return
// 	}

// 	// Panggil service untuk mengambil pengguna berdasarkan ID
// 	userResponse, err := h.userService.FindUserByID(uint(userID))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": userResponse,
// 	})
// }

func (h *userHandler) Login(c *gin.Context) {
	var loginRequest user.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s,condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}
	tokenString, err := h.userService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"token": tokenString,
		},
	})
}

// func (h *userHandler) DeleteUser(c *gin.Context) {
// 	// Mendapatkan ID pengguna dari URL atau permintaan
// 	userIDStr := c.Param("id")
// 	userID, err := strconv.ParseUint(userIDStr, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": "Invalid user ID",
// 		})
// 		return
// 	}

// 	// Menggunakan service untuk menghapus pengguna
// 	err = h.userService.DeleteUser(uint(userID))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "User deleted successfully",
// 	})
// }

func (h *userHandler) DeleteUser(c *gin.Context) {
	// Mendapatkan ID pengguna dari URL atau permintaan
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid user ID",
		})
		return
	}

	// Menggunakan service untuk menghapus pengguna
	err = h.userService.DeleteUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
