package ports

import "notes/internal/core/notes"

type NotesPort interface {
	CreateNote(note *notes.Note) error
	GetNotes(id string) ([]*notes.Note, error)
}