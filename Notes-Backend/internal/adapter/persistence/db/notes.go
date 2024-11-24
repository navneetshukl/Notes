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
func (nr *NotesRepo) GetNotes() ([]*notes.Note, error) {

	data := []*notes.Note{}
	err := nr.db.Find(&data).Error
	if err != nil {
		nr.log.Error("Failed to get notes", err)
		return nil, ErrGettingNotes
	}
	return data, nil
}
