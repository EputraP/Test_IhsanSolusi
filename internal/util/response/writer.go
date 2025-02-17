package response

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

func JSON(ctx *fiber.Ctx, statusCode int, message string, data interface{}) {
	resp := NewResponse(statusCode, message, data)

	ctx.Status(statusCode).JSON(resp)
}

func Error(ctx *fiber.Ctx, statusCode int, message string) {
	JSON(ctx, statusCode, message, nil)
	ctx.SendStatus(400)
}

func ValidationError(ctx *fiber.Ctx, errs validator.ValidationErrors) {
	err := errs[0]

	msg := fmt.Sprintf("validation error: field: [%v] rule: [%v] ", err.Field(), err.ActualTag())

	if err.Param() != "" {
		msg += fmt.Sprintf("rule value: [%v] ", err.Param())
	}

	msg += fmt.Sprintf("actual value: [%v]", err.Value())

	Error(ctx, http.StatusBadRequest, msg)
}

func UnknownError(ctx *fiber.Ctx, err error) {
	log.Println(err)
	Error(ctx, http.StatusInternalServerError, "Something went wrong")
}

func FromRequest(ctx *fiber.Ctx, response *http.Response, message string, data interface{}) {
	for key := range response.Header {
		ctx.Set(key, response.Header.Get(key))
		ctx.Status(http.StatusOK)
	}
}
