package handler

import (
	"company-project/employee"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type employeeHandler struct {
	employeeService employee.Service
}

func NewEmployeeHandler(service employee.Service) *employeeHandler {
	return &employeeHandler{service}
}

func (h *employeeHandler) GetAllEmployee(c *gin.Context) {
	empResponses, err := h.employeeService.FindAllEmployees()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": empResponses,
	})
}

func (h *userHandler) GetUserById(c *gin.Context) {
	userIDParam := c.Param("id")

	// Ubah ID pengguna menjadi tipe data yang sesuai (misalnya uint)
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid user ID",
		})
		return
	}

	// Panggil service untuk mengambil pengguna berdasarkan ID
	userResponse, err := h.userService.FindUserByID(uint(userID))
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
