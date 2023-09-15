package employee

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Employee, error)
	FindByID(ID uint) (Employee, error)
	Create(employee Employee) (Employee, error)
	Update(Employee Employee) (Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Employee, error) {
	var employees []Employee
	err := r.db.Preload("Employee").Find(&employees).Error
	return employees, err
}

func (r *repository) FindByID(ID uint) (Employee, error) {
	var employee Employee
	err := r.db.First(&employee, ID).Error
	return employee, err
}

func (r *repository) Create(employee Employee) (Employee, error) {
	err := r.db.Create(&employee).Error
	return employee, err
}

func (r *repository) Update(employee Employee) (Employee, error) {
	err := r.db.Save(&employee).Error
	return employee, err
}
