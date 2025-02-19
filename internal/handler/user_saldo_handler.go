package handler

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
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

	logger.Info("Received new request in TabungHandler", "method", c.Method(), "path", c.Path())

	var transactionBody *dto.TransactionBody

	if err := c.BodyParser(&transactionBody); err != nil {
		logger.Error("Failed to parse request body in TabungHandler", "error", err)
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Debug("Parsed request body in TabungHandler", "transactionBody", *transactionBody)

	if noRekErr := validator.Validate12DigitNumber(transactionBody.NoRekening); noRekErr != nil {
		logger.Warn("Invalid rekening number", "noRekening", transactionBody.NoRekening, "error", noRekErr.Error())
		response.Error(c, 400, noRekErr.Error())
		return
	}

	if nominalErr := validator.ValidateRupiahNominal(transactionBody.Nominal); nominalErr != nil {
		logger.Warn("Invalid nominal amount", "nominal", transactionBody.Nominal, "error", nominalErr.Error())
		response.Error(c, 400, nominalErr.Error())
		return
	}

	resp, err := h.userSaldoService.TabungTarikSaldo("tabung", &dto.TransactionBody{NoRekening: transactionBody.NoRekening, Nominal: transactionBody.Nominal})
	if err != nil {
		logger.Error("Failed to process transaction in TabungHandler", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("Tabung transaction successful", "resp", *resp)
	response.JSON(c, 200, "Tabung transaction Success", resp)
}

func (h UserSaldoHandler) TarikHandler(c *fiber.Ctx) {

	var transactionBody *dto.TransactionBody

	if err := c.BodyParser(&transactionBody); err != nil {
		logger.Error("Failed to parse request body in TarikHandler", "error", err)
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	if noRekErr := validator.Validate12DigitNumber(transactionBody.NoRekening); noRekErr != nil {
		logger.Warn("Invalid rekening number", "noRekening", transactionBody.NoRekening, "error", noRekErr.Error())
		response.Error(c, 400, noRekErr.Error())
		return
	}

	if nominalErr := validator.ValidateRupiahNominal(transactionBody.Nominal); nominalErr != nil {
		logger.Warn("Invalid nominal amount", "nominal", transactionBody.Nominal, "error", nominalErr.Error())
		response.Error(c, 400, nominalErr.Error())
		return
	}

	resp, err := h.userSaldoService.TabungTarikSaldo("tarik", &dto.TransactionBody{NoRekening: transactionBody.NoRekening, Nominal: transactionBody.Nominal})
	if err != nil {
		logger.Error("Failed to process transaction in TarikHandler", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("Tarik transaction successful", "resp", *resp)
	response.JSON(c, 200, "Tarik Success", resp)
}

func (h UserSaldoHandler) GetSaldoByNoRekHandler(c *fiber.Ctx) {

	noRek := c.Params("no_rekening")

	resp, err := h.userSaldoService.GetUserSaldoByNoRek(&noRek)
	if err != nil {
		logger.Error("Failed to get saldo by rekening in GetSaldoByNoRekHandler", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("Successfully retrieved saldo for rekening", "resp", *resp)
	response.JSON(c, 200, "Getting User Saldo Success", resp)
}
