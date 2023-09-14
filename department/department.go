package department

import (
	"company-project/employee"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name     string
	Employee []employee.Employee
}
