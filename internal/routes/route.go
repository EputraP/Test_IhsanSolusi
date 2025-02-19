package routes

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/handler"
	"github.com/gofiber/fiber"
)

type Handlers struct {
	UserHandler      *handler.UserHandler
	UserSaldoHandler *handler.UserSaldoHandler
}

type Middlewares struct {
	UserSaldoMiddleware fiber.Handler
}

func Build(srv *fiber.App, h Handlers, middleware Middlewares) {

	srv.Post("/daftar", h.UserHandler.UserHandler)
	srv.Post("/tabung", h.UserSaldoHandler.TabungHandler)
	srv.Post("/tarik", h.UserSaldoHandler.TarikHandler)
	srv.Get("/saldo/:no_rekening", middleware.UserSaldoMiddleware, h.UserSaldoHandler.GetSaldoByNoRekHandler)
}
