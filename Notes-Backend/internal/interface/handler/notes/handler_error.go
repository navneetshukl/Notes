package notes

import "github.com/gofiber/fiber/v2"

type HandlerError struct {
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

func handlerError(ctx *fiber.Ctx, err error) error {
	var errResponse HandlerError

	switch err {

	default:
		errResponse.StatusCode = fiber.StatusInternalServerError
		errResponse.Message = "something went wrong"
		errResponse.Data = map[string]interface{}{}

	}
	return ctx.JSON(errResponse)

}
