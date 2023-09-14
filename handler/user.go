package handler

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "company-project/user" // Sesuaikan dengan path yang benar
)

type userHandler struct {
    userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
    return &userHandler{service}
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
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