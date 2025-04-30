package repository

import (
	"github.com/dzaakk/simple-grpc-app/internal/user/model"
)

type UserRepository interface {
	Create(user model.TUser) error
	GetByID(id int) (*model.TUser, error)
	Update(user model.TUser) error
	Delete(id int) error
	GetAll() ([]*model.TUser, error)
}
