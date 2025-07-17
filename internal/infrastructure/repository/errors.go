package repository

import "errors"

var (
	NotFound     = errors.New("task not found")
	LoginExist   = errors.New("the user with this username already exists")
	Unauthorized = errors.New("session is not found")
)
