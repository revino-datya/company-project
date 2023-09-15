package department

import (
	"company-project/employee"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name     string `gorm:"size:20"`
	Employee []employee.Employee
}
