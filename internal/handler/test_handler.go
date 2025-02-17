package handler

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/response"
	"github.com/gofiber/fiber"
)

type TestHandler struct {
	testService service.TestService
}

type TestHandlerConfig struct {
	TestService service.TestService
}

func NewTestHandler(config TestHandlerConfig) *TestHandler {
	return &TestHandler{
		testService: config.TestService,
	}
}

func (h TestHandler) TestHandler(c *fiber.Ctx) {

	resp, _ := h.testService.TestService()

	response.JSON(c, 201, "Test Handler Succes", resp)
}
