package apperrors

import "errors"

var ErrUserNotFound = errors.New("user not found")

var ErrInvalidCredentials = errors.New("invalid credentials")

var ErrEmailAlreadyExists = errors.New("user with this email already exists")
