package usecase

import "github.com/dzaakk/simple-grpc-app/internal/user/model"

type UserUseCase interface {
	Login(email, password string) (*model.User, error)
	CreateUser(user model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]*model.User, error)
}
