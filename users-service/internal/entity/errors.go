package entity

import "errors"

var (
	ErrUserNotFound      = errors.New("user does not exist")
	ErrUserAlreadyExists = errors.New("user with such email already exists")
)
