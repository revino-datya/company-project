package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/go-playground/validator/v10"

	// "encoding/json"

	// "fmt"

	"company-project/user"

	"strconv"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

// func (h *bookHandler) PostBooksHandler(c *gin.Context) {
// 	var bookRequest book.BookRequest

// 	err := c.ShouldBindJSON(&bookRequest)

// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			errorMessages := []string{}
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
// 				errorMessages = append(errorMessages, errorMessage)
// 			}
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": errorMessages,
// 			})
// 			return
// 		case *json.UnmarshalTypeError:
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	jwtClaims, _ := c.Get("jwtClaims")
// 	claims, _ := jwtClaims.(jwt.MapClaims)
// 	userID, _ := claims["sub"].(float64)

// 	book, err := h.bookService.Create(bookRequest, uint(userID))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": book,
// 	})
// }

// func (h *userHandler) GetAllUsers(c *gin.Context) {
// 	jwtClaims, _ := c.Get("jwtClaims")
// 	claims, _ := jwtClaims.(jwt.MapClaims)
// 	userID, _ := claims["sub"].(float64)

// 	users, err := h.userService.FindAll(uint(userID))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}

// 	var usersResponse []user.UserResponse

// 	for _, u := range users {
// 		userResponse := user.ConvertToUserResponse(u)
// 		usersResponse = append(usersResponse, userResponse)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": usersResponse,
// 	})
// }

func (h *userHandler) GetUserById(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	u, err := h.userService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := user.ConvertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

// func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
// 	var bookRequest book.BookRequest

// 	err := c.ShouldBindJSON(&bookRequest)

// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			errorMessages := []string{}
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
// 				errorMessages = append(errorMessages, errorMessage)
// 			}
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": errorMessages,
// 			})
// 			return
// 		case *json.UnmarshalTypeError:
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	ID, _ := strconv.Atoi(c.Param("id"))
// 	b, err := h.bookService.Update(ID, bookRequest)
// 	bookResponse := book.ConvertToBookResponse(b)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": bookResponse,
// 	})
// }

// func (h *bookHandler) DeleteBook(c *gin.Context) {
// 	ID, _ := strconv.Atoi(c.Param("id"))
// 	b, err := h.bookService.Delete(ID)
// 	bookResponse := book.ConvertToBookResponse(b)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": bookResponse,
// 	})
// }
