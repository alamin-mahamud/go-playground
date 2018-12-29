package entity

import "errors"

var (
	ErrNotFound        = errors.New("Not Found")
	ErrCannotBeDeleted = errors.New("Can't Be deleted")
)
