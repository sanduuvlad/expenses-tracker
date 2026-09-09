package apperrors

import "errors"

var ErrUserNotFound = errors.New("user not found")

var ErrInvalidCredentials = errors.New("invalid credentials")
