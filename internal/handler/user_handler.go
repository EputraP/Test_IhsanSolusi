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

	logger.Info("Received new request in UserHandler", "method", c.Method(), "path", c.Path())

	var createUserBody *dto.CreateUserBody

	if err := c.BodyParser(&createUserBody); err != nil {
		logger.Error("Failed to parse request body in UserHandler", "error", err)
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}
	logger.Debug("Parsed request body in UserHandler", "createUserBody", *createUserBody)

	if nameErr := validator.ValidateName(createUserBody.Nama); nameErr != nil {
		logger.Warn("Invalid name provided in UserHandler", "name", createUserBody.Nama, "error", nameErr.Error())
		response.Error(c, 400, nameErr.Error())
		return
	}
	if noHPErr := validator.ValidatePhoneNumber(createUserBody.NoHP); noHPErr != nil {
		logger.Warn("Invalid phone number in UserHandler", "phone", createUserBody.NoHP, "error", noHPErr.Error())
		response.Error(c, 400, noHPErr.Error())
		return
	}
	if nikErr := validator.ValidateNIK(createUserBody.NIK); nikErr != nil {
		logger.Warn("Invalid NIK provided in UserHandler", "nik", createUserBody.NIK, "error", nikErr.Error())
		response.Error(c, 400, nikErr.Error())
		return
	}

	resp, err := h.userService.CreateUser(createUserBody)
	if err != nil {
		logger.Error("Failed to create user in UserHandler", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("User created successfully in UserHandler", "resp", *resp)

	response.JSON(c, 200, "Creating User Succes", resp)
}
