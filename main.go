package main

import (
	// "company-project/handler"
	"company-project/department"
	"company-project/handler"
	"company-project/initializers"

	// "company-project/user"
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

	router := gin.Default()

	routerV1 := router.Group("/v1")

	routerV1.POST("/department", departmentHandler.PostBooksHandler)

	routerV1.PUT("/department/:id", departmentHandler.UpdateDepartmentHandler)

	routerV1.DELETE("/department/:id", departmentHandler.DeleteDepartment)

	routerV1.GET("/department", departmentHandler.GetAllDepartments)

	routerV1.GET("/department/:id", departmentHandler.GetDepartmentByID)

	// routerV1.GET("/user/:id", userHandler.GetUserById)

	// routerV1.POST("/signup", userHandler.Signup)

	// routerV1.POST("/login", userHandler.Login)

	// routerV1Books := routerV1.Group("/book", middleware.RequireAuth)

	// routerV1Books.POST("/books", bookHandler.PostBooksHandler)

	// routerV1Books.GET("/books", bookHandler.GetAllBooks)

	// routerV1Books.GET("/books/:id", bookHandler.GetBookById)

	// routerV1Books.PUT("/books/:id", bookHandler.UpdateBookHandler)

	// routerV1Books.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":3030")
}
