package notes

import "errors"

var (
	ErrTitleEmpty = errors.New("title cannot be empty")
	ErrContentEmpty = errors.New("content cannot be empty")
)
