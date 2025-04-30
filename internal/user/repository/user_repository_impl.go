package repository

import (
	"database/sql"

	"github.com/dzaakk/simple-grpc-app/internal/user/model"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// Create implements UserRepository.
func (u *userRepositoryImpl) Create(user model.TUser) error {
	panic("unimplemented")
}

// Delete implements UserRepository.
func (u *userRepositoryImpl) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements UserRepository.
func (u *userRepositoryImpl) GetAll() ([]*model.TUser, error) {
	panic("unimplemented")
}

// GetByID implements UserRepository.
func (u *userRepositoryImpl) GetByID(id int) (*model.TUser, error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (u *userRepositoryImpl) Update(user model.TUser) error {
	panic("unimplemented")
}
