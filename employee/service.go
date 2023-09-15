package employee

type Service interface {
	FindAllEmployees() ([]Employee, error)
	FindEmployeeByID(ID uint) (Employee, error)
	// UpdateEmployee(ID uint, updateEmpRequest UpdateEmpRequest) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindAllEmployees() ([]Employee, error) {
	employees, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	empResponses := make([]Employee, len(employees))

	return empResponses, nil
}

func (s *service) FindEmployeeByID(employeeID uint) (Employee, error) {
	employee, err := s.repository.FindByID(employeeID)
	if err != nil {
		return Employee{}, err
	}

	// Konversi entitas User ke UserResponse menggunakan mapper
	// userResponse := ConvertToUserResponse(user)

	return employee, nil
}

// func (s *service) Update(ID uint, updateEmpRequest UpdateEmpRequest) (Employee, error) {
// 	employee, err := s.repository.FindByID(ID)
// 	if updateEmpRequest.Name != "" {
// 		employee.Name = updateEmpRequest.Name
// 	}
// 	if updateEmpRequest.Phone != 0 {
// 		employee.Phone = updateEmpRequest.Phone
// 	}
// 	if updateEmpRequest.Department_Id != 0 {
// 		// employee.Department_Id = updateEmpRequest.Department_Id
// 	}
// 	newEmployee, err := s.repository.Update(employee)
// 	return newEmployee, err
// }
