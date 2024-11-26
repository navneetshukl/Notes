package notes

import (
	"errors"
	"net/http"
	"notes/internal/core/notes"

	"github.com/gofiber/fiber/v2"
)

type NotesHandler struct {
	notesUseCase notes.NotesUsercase
}

func NewNotesHandler(notesUseCase notes.NotesUsercase) *NotesHandler {
	return &NotesHandler{notesUseCase: notesUseCase}
}
func (nu *NotesHandler) GetNote(ctx *fiber.Ctx) error {

	reqBody := notes.NoteReq{}
	err := ctx.BodyParser(&reqBody)
	if err != nil {
		return handlerError(ctx, errors.New("something went wrong"))
	}

	err = nu.notesUseCase.CreateNote(reqBody)
	if err != nil {
		return handlerError(ctx, err)
	}

	return ctx.JSON(fiber.Map{
		"status_code": http.StatusOK,
		"message":     "note created successfully",
		"data":        nil,
	})

}
