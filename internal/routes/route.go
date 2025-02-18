package routes

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/handler"
	"github.com/gofiber/fiber"
)

type Handlers struct {
	TestHandler      *handler.TestHandler
	UserHandler      *handler.UserHandler
	UserSaldoHandler *handler.UserSaldoHandler
}

func Build(srv *fiber.App, h Handlers) {

	srv.Post("/daftar", h.UserHandler.UserHandler)

}
