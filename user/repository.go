package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create() (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	return user, err
}
func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
