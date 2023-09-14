package department

func ConvertToDepartmentResponse(d Department) DepartmentResponse {
	return DepartmentResponse{
		ID:   int(d.ID),
		Name: d.Name,
	}
}
