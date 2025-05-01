package usecase

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

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

func (u *userUseCaseImpl) CreateUser(user model.User) error {
	dbUser := model.TUser{
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		RoleID:     user.RoleID,
		RoleName:   user.RoleName,
		LastAccess: sql.NullTime{Time: time.Now(), Valid: true},
	}
	return u.repo.Create(dbUser)
}

func (u *userUseCaseImpl) DeleteUser(id int) error {
	return u.repo.Delete(id)
}

func (u *userUseCaseImpl) GetAllUsers() ([]*model.User, error) {
	listUser, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, user := range listUser {
		users = append(users, &model.User{
			Name:       user.Name,
			Email:      user.Email,
			RoleID:     user.RoleID,
			RoleName:   user.RoleName,
			LastAccess: user.LastAccess.Time.String(),
		})
	}

	return users, nil
}

func (u *userUseCaseImpl) GetUserByEmail(email string) (*model.User, error) {
	dbUser, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}

	return &model.User{
		Name:       dbUser.Name,
		Email:      dbUser.Email,
		RoleID:     dbUser.RoleID,
		RoleName:   dbUser.RoleName,
		LastAccess: dbUser.LastAccess.Time.String(),
	}, nil
}

// GetUserByID implements UserUseCase.
func (u *userUseCaseImpl) GetUserByID(id int) (*model.User, error) {
	dbUser, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}

	return &model.User{
		Name:       dbUser.Name,
		Email:      dbUser.Email,
		RoleID:     dbUser.RoleID,
		RoleName:   dbUser.RoleName,
		LastAccess: dbUser.LastAccess.Time.String(),
	}, nil
}

func (u *userUseCaseImpl) Login(email string, password string) (*model.User, error) {
	dbUser, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if dbUser.Password != password {
		return nil, fmt.Errorf("invalid password")
	}

	return &model.User{
		Name:       dbUser.Name,
		Email:      dbUser.Email,
		RoleID:     dbUser.RoleID,
		RoleName:   dbUser.RoleName,
		LastAccess: dbUser.LastAccess.Time.String(),
	}, nil
}

func (u *userUseCaseImpl) UpdateUser(user model.User) error {
	userID, _ := strconv.Atoi(user.ID)
	dbUser := model.TUser{
		ID:         userID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		RoleID:     user.RoleID,
		RoleName:   user.RoleName,
		LastAccess: sql.NullTime{Time: time.Now(), Valid: true},
	}
	return u.repo.Update(dbUser)
}
