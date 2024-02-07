package entity

import (
	"errors" // define error
)

// Define new error
var (
	ErrPswdCnntHash = errors.New("Password can not be hashed!")
	ErrExtUsername = errors.New("Username exitsted!")
	ErrNotFoundUsername = errors.New("Username not found!")
	ErrPswdNotMatch = errors.New("Password does not match!")
	ErrDelAdmin = errors.New("Admin account can not be deleted!")
)
