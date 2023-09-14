package department

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Department, error)
	FindByID(ID int) (Department, error)
	Create(department Department) (Department, error)
	Update(department Department) (Department, error)
	Delete(department Department) (Department, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Department, error) {
	var departments []Department
	err := r.db.Find(&departments).Error
	return departments, err
}

func (r *repository) FindByID(ID int) (Department, error) {
	var department Department
	err := r.db.First(&department, ID).Error
	return department, err
}
func (r *repository) Create(department Department) (Department, error) {
	err := r.db.Create(&department).Error
	return department, err
}

func (r *repository) Update(department Department) (Department, error) {
	err := r.db.Save(&department).Error
	return department, err
}

func (r *repository) Delete(department Department) (Department, error) {
	err := r.db.Delete(&department).Error
	return department, err
}
