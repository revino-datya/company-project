package initializers

import (
	"company-project/department"
	"company-project/employee"
	"company-project/user"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(department.Department{}, user.User{}, employee.Employee{})
	return err
}
