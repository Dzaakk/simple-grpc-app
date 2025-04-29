package model

import "database/sql"

type TUser struct {
	ID         int
	Name       string
	Email      string
	Password   string
	RoleID     string
	RoleName   string
	LastAccess sql.NullTime
}

type User struct {
	ID         string
	Name       string
	Email      string
	Password   string
	RoleID     string
	RoleName   string
	LastAccess string
}
