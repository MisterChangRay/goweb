package user

import (
	"time"
)

type UserBase struct {
	Name     string
	Account  string
	Password string
}

type UserAdd struct {
	UserBase
	CreateTime time.Time
}

type UserUpdate struct {
	UserBase
	Id         int
	UpdateTime time.Time
}

type UserList struct {
	UserBase
	UserAdd
}
