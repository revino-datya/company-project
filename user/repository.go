package user

import (
	"gorm.io/gorm"
)

// Repository adalah interface yang mendefinisikan operasi CRUD untuk entitas User.
type Repository interface {
	Create(user User) (User, error)
	FindAll() ([]User, error)
	FindUserByIDWithEmployee(ID uint) (User, error)
	Update(user User) (User, error)
	Delete(user User) error
	FindByEmail(email string) (User, error)
}

// repository adalah implementasi dari Repository.
type repository struct {
	db *gorm.DB
}

// NewRepository adalah fungsi pembuat yang digunakan untuk membuat instance Repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// Create digunakan untuk membuat entitas User baru.
func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// FindAll digunakan untuk mendapatkan semua entitas User dari database.
func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindUserByIDWithEmployee(userID uint) (User, error) {
    var user User
    err := r.db.Preload("Employee").First(&user, userID).Error
    return user, err
}

// Update digunakan untuk memperbarui entitas User yang ada.
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

// Delete digunakan untuk menghapus entitas User yang ada.
func (r *repository) Delete(user User) error {
	err := r.db.Delete(&user).Error
	return err
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}
