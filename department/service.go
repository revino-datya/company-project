package department

type Service interface {
	FindAll() ([]Department, error)
	FindByID(ID int) (Department, error)
	Create(departmentRequest DepartmentRequest) (Department, error)
	Update(ID int, departmentRequest DepartmentRequest) (Department, error)
	Delete(ID int) (Department, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) FindAll() ([]Department, error) {
	departments, err := s.repository.FindAll()
	return departments, err
}

func (s *service) FindByID(ID int) (Department, error) {
	department, err := s.repository.FindByID(ID)
	return department, err
}
func (s *service) Create(departmentRequest DepartmentRequest) (Department, error) {
	department := Department{
		Name: departmentRequest.Name,
	}
	newDepartment, err := s.repository.Create(department)
	return newDepartment, err
}

func (s *service) Update(ID int, departmentRequest DepartmentRequest) (Department, error) {
	department, err := s.repository.FindByID(ID)
	if departmentRequest.Name != "" {
		department.Name = departmentRequest.Name
	}
	newDepartment, err := s.repository.Update(department)
	return newDepartment, err
}

func (s *service) Delete(ID int) (Department, error) {
	department, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(department)
	return department, err
}
