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

func (r *userRepositoryImpl) Create(user model.TUser) error {
	query := `INSERT INTO users (name, email, password, role_id, role_name last_access)
			  VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.RoleName, user.LastAccess)
	return err
}

func (r *userRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *userRepositoryImpl) GetAll() ([]*model.TUser, error) {
	query := `SELECT id, username, email, password, role_id, role_name, last_access FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.TUser
	for rows.Next() {
		var user model.TUser
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.RoleName, &user.LastAccess); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userRepositoryImpl) GetByID(id int) (*model.TUser, error) {
	query := `SELECT id, name, email, password, role_id, role_name, last_access FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user model.TUser
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.RoleName, &user.LastAccess)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByEmail(email string) (*model.TUser, error) {
	query := `SELECT id, name, email, password, role_id, role_name, last_access FROM users WHERE email= $1`
	row := r.db.QueryRow(query, email)

	var user model.TUser
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.RoleName, &user.LastAccess)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Update(user model.TUser) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role_id = $4, role_name = $5, last_access = $6
			  WHERE id = $7`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.RoleName, user.LastAccess, user.ID)
	return err
}
