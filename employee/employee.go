package employee

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string
	Phone        int
	UserID       uint `gorm:"unique"`
	DepartmentID uint
}
