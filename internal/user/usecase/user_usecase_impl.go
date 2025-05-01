package usecase

import (
	"github.com/dzaakk/simple-grpc-app/internal/user/model"
	"github.com/dzaakk/simple-grpc-app/internal/user/repository"
)

type userUseCaseImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{
		repo: repo,
	}
}

// CreateUser implements UserUseCase.
func (u *userUseCaseImpl) CreateUser(user model.User) error {
	panic("unimplemented")
}

// DeleteUser implements UserUseCase.
func (u *userUseCaseImpl) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetAllUsers implements UserUseCase.
func (u *userUseCaseImpl) GetAllUsers() ([]*model.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements UserUseCase.
func (u *userUseCaseImpl) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}

// GetUserByID implements UserUseCase.
func (u *userUseCaseImpl) GetUserByID(id int) (*model.User, error) {
	panic("unimplemented")
}

// Login implements UserUseCase.
func (u *userUseCaseImpl) Login(email string, password string) (*model.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserUseCase.
func (u *userUseCaseImpl) UpdateUser(user model.User) error {
	panic("unimplemented")
}
