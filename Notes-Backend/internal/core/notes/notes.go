package notes

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	ID      string `gorm:"column:id" json:"id"`
}
type NoteReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NotesUsercase interface {
	CreateNote(note NoteReq) (error)
	GetNotes(id string) ([]*Note, error)
}
