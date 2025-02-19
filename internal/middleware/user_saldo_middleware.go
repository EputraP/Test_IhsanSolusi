package middleware

import (
	"fmt"

	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/response"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/validator"
	"github.com/gofiber/fiber"
)

func ValidateNoRekening() fiber.Handler {
	return func(c *fiber.Ctx) {

		noRek := c.Params("no_rekening")

		fmt.Println("noRek: ", noRek)

		if noRekErr := validator.Validate12DigitNumber(noRek); noRekErr != nil {
			logger.Warn("Invalid rekening number", "noRekening", noRek, "error", noRekErr.Error())
			response.Error(c, 400, noRekErr.Error())
			return
		}

		c.Next()
	}
}
