package storage

import "fmt"

var (
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrUserExists   = fmt.Errorf("user already exists")
	ErrAppNotFound  = fmt.Errorf("app not found")
)
