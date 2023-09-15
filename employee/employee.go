package employee

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string `gorm:"size:35"`
	Phone        int
	UserID       uint `gorm:"unique"`
	DepartmentID *uint
}
