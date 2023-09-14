package main

import (
	// "company-project/handler"
	"company-project/department"
	"company-project/employee"
	"company-project/handler"
	"company-project/initializers"

	"company-project/user"
	// "company-project/employee"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initializers.LoadEnvVariables()

	db, err = initializers.ConnectToDatabase()
	err = initializers.SyncDatabase(db)
	if err != nil {
		log.Fatal("Connection to database failed")
	}
}

func main() {
	departmentRepository := department.NewRepository(db)
	departmentService := department.NewService(departmentRepository)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	employeeRepository := employee.NewRepository(db)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository,employeeRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	routerV1.POST("/department", departmentHandler.PostDepartmentHandler)

	routerV1.PUT("/department/:id", departmentHandler.UpdateDepartmentHandler)

	routerV1.DELETE("/department/:id", departmentHandler.DeleteDepartment)

	routerV1.GET("/department", departmentHandler.GetAllDepartments)

	routerV1.GET("/department/:id", departmentHandler.GetDepartmentByID)


	routerV1.POST("/signup", userHandler.CreateUserHandler)

	// routerV1.POST("/login", userHandler.Login)

	router.Run(":3030")
}
