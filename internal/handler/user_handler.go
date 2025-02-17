package handler

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/response"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/validator"
	"github.com/gofiber/fiber"
)

type UserHandler struct {
	userService service.UserService
}

type UserHandlerConfig struct {
	UserService service.UserService
}

func NewUserHandler(config UserHandlerConfig) *UserHandler {
	return &UserHandler{
		userService: config.UserService,
	}
}

func (h UserHandler) UserHandler(c *fiber.Ctx) {

	var createUserBody *dto.CreateUserBody

	if err := c.BodyParser(&createUserBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	if nameErr := validator.ValidateName(createUserBody.Nama); nameErr != nil {
		response.Error(c, 400, nameErr.Error())
		return
	}
	if noHPErr := validator.ValidatePhoneNumber(createUserBody.NoHP); noHPErr != nil {
		response.Error(c, 400, noHPErr.Error())
		return
	}
	if nikErr := validator.ValidateNIK(createUserBody.NIK); nikErr != nil {
		response.Error(c, 400, nikErr.Error())
		return
	}

	resp, err := h.userService.CreateUser(createUserBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Creating User Succes", resp)
}
