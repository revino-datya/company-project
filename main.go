package main

import (
	// "company-project/handler"
	"company-project/department"
	"company-project/handler"
	"company-project/initializers"
	"company-project/middleware"

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

	// employeeRepository := employee.NewRepository(db)

	userRepository := user.NewRepository(db)
	// userService := user.NewService(userRepository, employeeRepository)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	routerV1User := routerV1.Group("", middleware.RequireAuth)

	routerV1User.POST("/department", departmentHandler.PostDepartmentHandler)

	routerV1User.PUT("/department/:id", departmentHandler.UpdateDepartmentHandler)

	routerV1User.DELETE("/department/:id", departmentHandler.DeleteDepartment)

	routerV1User.GET("/department", departmentHandler.GetAllDepartments)

	routerV1User.GET("/department/:id", departmentHandler.GetDepartmentByID)

	routerV1.GET("/user", userHandler.GetAllUser)
	routerV1.GET("/user/:id", userHandler.GetUserById)

	routerV1.PUT("/user/:id", userHandler.UpdateUser)
	routerV1.DELETE("/user/:id", userHandler.DeleteUser)

	routerV1.POST("/signup", userHandler.CreateUser)
	routerV1.POST("/login", userHandler.Login)

	router.Run(":3030")
}
