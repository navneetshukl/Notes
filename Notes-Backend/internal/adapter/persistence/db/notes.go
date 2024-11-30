package db

import (
	"notes/internal/core/notes"
	"notes/pkg/logging"

	"gorm.io/gorm"
)

type NotesRepo struct {
	db  *gorm.DB
	log *logging.Log
}

func NewNotesRepo(db *gorm.DB, log *logging.Log) *NotesRepo {
	return &NotesRepo{db: db, log: log}
}
func (nr *NotesRepo) CreateNote(note *notes.Note) error {
	err := nr.db.Create(note).Error
	if err != nil {
		nr.log.Error("Failed to add note to DB", err)
		return ErrAddingNote
	}
	return nil

}
func (nr *NotesRepo) GetNotes(id string) ([]*notes.Note, error) {
	// Initialize the result slice
	data := []*notes.Note{}

	// Query the database filtering by userID
	err := nr.db.Where("user_id = ?", id).Find(&data).Error
	if err != nil {
		// Log error if something goes wrong
		nr.log.Error("Failed to get notes by userID", err)
		return nil, ErrGettingNotes
	}

	// Return the retrieved notes
	return data, nil
}
