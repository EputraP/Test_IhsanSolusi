package routes

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/handler"
	"github.com/gofiber/fiber"
)

type Handlers struct {
	TestHandler *handler.TestHandler
}

func Build(srv *fiber.App, h Handlers) {
	test := srv.Group("/test")
	test.Get("/hello", h.TestHandler.TestHandler)
}
