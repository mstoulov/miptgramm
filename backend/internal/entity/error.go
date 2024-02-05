package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidUserID = errors.New("invalid user id")
var ErrInvalidUserName = errors.New("invalid user name")
var ErrInvalidPassword = errors.New("invalid password")
var ErrUserExists = errors.New("user with this login exists")
var ErrIncorrectPassword = errors.New("password is not correct")
var ErrInvalidRequest = errors.New("invalid request")
var ErrInvalidRequestType = errors.New("invalid request type")
var ErrPermissionDenied = errors.New("permission denied")
