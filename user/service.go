package user

type Service interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(userRequest UserRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	return user, err
}

// func (s *service) Create(userRequest UserRequest) (User, error) {
// 	user := User{
// 		Email:    userRequest.Email,
// 		Password: userRequest.Password,
// 	}

// 	newUser, err := s.repository.Create(user)
// 	return newUser, err
// }
