package user

type UserService struct {
	userEntity *UserEntity `inject:""`
}

func (s *UserService) Create() *User {
	return &User{
		Name: "Awesome Goose",
	}
}

func (s *UserService) List() *User {
	return &User{
		Name: "Awesome Goose",
	}
}

func (s *UserService) Get() *User {
	return &User{
		Name: "Awesome Goose",
	}
}

func (s *UserService) Update() *User {
	return &User{
		Name: "Awesome Goose",
	}
}

func (s *UserService) Delete() *User {
	return &User{
		Name: "Awesome Goose",
	}
}
