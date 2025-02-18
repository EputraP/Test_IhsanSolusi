package handler

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/response"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/validator"
	"github.com/gofiber/fiber"
)

type UserSaldoHandler struct {
	userSaldoService service.UserSaldoService
}

type UserSaldoHandlerConfig struct {
	UserSaldoService service.UserSaldoService
}

func NewUserSaldoHandler(config UserSaldoHandlerConfig) *UserSaldoHandler {
	return &UserSaldoHandler{
		userSaldoService: config.UserSaldoService,
	}
}

func (h UserSaldoHandler) TabungHandler(c *fiber.Ctx) {

	var transactionBody *dto.TransactionBody

	if err := c.BodyParser(&transactionBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	if noRekErr := validator.Validate12DigitNumber(transactionBody.NoRekening); noRekErr != nil {
		response.Error(c, 400, noRekErr.Error())
		return
	}
	if nominalErr := validator.ValidateRupiahNominal(transactionBody.Nominal); nominalErr != nil {
		response.Error(c, 400, nominalErr.Error())
		return
	}

	// resp, err := h.userService.CreateUser(createUserBody)
	// if err != nil {
	// 	response.Error(c, 400, err.Error())
	// 	return
	// }
	response.JSON(c, 200, "Creating User Succes", "")
}

func (h UserSaldoHandler) GetSaldoByNoRekHandler(c *fiber.Ctx) {

	noRek := c.Params("no_rekening")

	if noRekErr := validator.Validate12DigitNumber(noRek); noRekErr != nil {
		response.Error(c, 400, noRekErr.Error())
		return
	}

	resp, err := h.userSaldoService.GetUserSaldoByNoRek(&noRek)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Creating User Succes", resp)
}
