package department

type DepartmentRequest struct {
	Name string `binding:"required"`
}

type UpdateDepartmentRequest struct {
	Name string `binding:"required"`
}
