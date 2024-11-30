package notes

import (
	"fmt"
	"notes/internal/adapter/persistence/db"
	"notes/internal/core/notes"
	"notes/pkg/logging"
)

type NotesUsecaseImpl struct {
	notesRepo *db.NotesRepo
	log       *logging.Log
}

func NewNotesUsecaseImpl(notesRepo *db.NotesRepo, log *logging.Log) *NotesUsecaseImpl {
	return &NotesUsecaseImpl{
		notesRepo: notesRepo,
		log:       log,
	}
}

// Add single note to DB
func (nu *NotesUsecaseImpl) CreateNote(note *notes.NoteReq) error {
	// check if any of the field is not empty
	if len(note.Title) == 0 || note.Title == "" {
		nu.log.Error("Title cannot be empty", fmt.Errorf("title cannot be empty"))
		return notes.ErrTitleEmpty
	}

	if len(note.Content) == 0 || note.Content == "" {
		nu.log.Error("Content cannot be empty", fmt.Errorf("content cannot be empty"))
		return notes.ErrContentEmpty
	}

	notesData := notes.Note{
		Title:   note.Title,
		Content: note.Content,
	}
	err := nu.notesRepo.CreateNote(&notesData)
	if err != nil {
		nu.log.Error("failed to add note to db", err)
		return db.ErrAddingNote
	}
	return nil
}

// Get all note of user from DB

func (nu *NotesUsecaseImpl) GetNotes(id string) ([]*notes.Note, error) {

	allNotes, err := nu.notesRepo.GetNotes(id)
	if err != nil {
		nu.log.Error("failed to get notes from db", err)
		return nil, db.ErrGettingNotes
	}
	return allNotes, nil

}
