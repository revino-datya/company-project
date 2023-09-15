package handler

import (
	"company-project/department"
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type departmentHandler struct {
	departmentService department.Service
}

func NewDepartmentHandler(service department.Service) *departmentHandler {
	return &departmentHandler{service}
}

func (h *departmentHandler) PostDepartmentHandler(c *gin.Context) {
	var departmentRequest department.DepartmentRequest

	err := c.ShouldBindJSON(&departmentRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
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

	// jwtClaims, _ := c.Get("jwtClaims")
	// claims, _ := jwtClaims.(jwt.MapClaims)
	// userID, _ := claims["sub"].(float64)

	department, err := h.departmentService.Create(departmentRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": department,
	})
}

func (h *departmentHandler) GetAllDepartments(c *gin.Context) {
	departments, err := h.departmentService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var departmentsResponse []department.DepartmentResponse
	for _, d := range departments {
		departmentResponse := department.ConvertToDepartmentResponse(d)
		departmentsResponse = append(departmentsResponse, departmentResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": departmentsResponse,
	})
}

func (h *departmentHandler) GetDepartmentByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	d, err := h.departmentService.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	departmentResponse := department.ConvertToDepartmentResponse(d)
	c.JSON(http.StatusBadRequest, gin.H{
		"data": departmentResponse,
	})
}

func (h *departmentHandler) UpdateDepartmentHandler(c *gin.Context) {
	var departmentRequest department.DepartmentRequest

	err := c.ShouldBindJSON(&departmentRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
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

	ID, _ := strconv.Atoi(c.Param("id"))
	d, err := h.departmentService.Update(ID, departmentRequest)
	departmentResponse := department.ConvertToDepartmentResponse(d)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": departmentResponse,
	})
}

func (h *departmentHandler) DeleteDepartment(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.departmentService.Delete(ID)
	departmentResponse := department.ConvertToDepartmentResponse(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": departmentResponse,
	})
}

// AMIN
