package db

import "errors"

var (
	ErrAddingNote   = errors.New("failed to add note")
	ErrGettingNotes = errors.New("failed to get notes")
)
