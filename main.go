package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EputraP/Test_IhsanSolusi/internal/constant"
	"github.com/EputraP/Test_IhsanSolusi/internal/handler"
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
	"github.com/EputraP/Test_IhsanSolusi/internal/routes"
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
	dbstore "github.com/EputraP/Test_IhsanSolusi/internal/store/db"
	"github.com/gofiber/fiber"
	"github.com/lpernett/godotenv"
)

func main() {
	env := os.Getenv(constant.EnvKeyEnv)

	if env != "prod" {
		err := godotenv.Load()

		if err != nil {
			log.Println("error loading env", err)
			log.Fatalln("error loading env", err)
		}
	}

	handlers := prepare()

	srv := fiber.New()

	routes.Build(srv, handlers)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := srv.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("Error running gin server: ", err)
		log.Fatalln("Error running gin server: ", err)
	}

}

func prepare() (handlers routes.Handlers) {

	db := dbstore.Get()

	testRepo := repository.NewTestRepository(db)
	userRepo := repository.NewUserRepository(db)
	userSaldoRepo := repository.NewUserSaldoRepository(db)

	testService := service.NewTestService(service.TestServiceConfig{
		TestRepo: testRepo,
	})
	userService := service.NewUserService(service.UserServiceConfig{
		UserRepo:      userRepo,
		UserSaldoRepo: userSaldoRepo,
	})
	userSaldoService := service.NewUserSaldoService(service.UserSaldoServiceConfig{
		UserSaldoRepo: userSaldoRepo,
	})

	testHandler := handler.NewTestHandler(handler.TestHandlerConfig{
		TestService: testService,
	})
	userHandler := handler.NewUserHandler(handler.UserHandlerConfig{
		UserService: userService,
	})
	userSaldoHandler := handler.NewUserSaldoHandler(handler.UserSaldoHandlerConfig{
		UserSaldoService: userSaldoService,
	})

	handlers = routes.Handlers{
		TestHandler:      testHandler,
		UserHandler:      userHandler,
		UserSaldoHandler: userSaldoHandler,
	}
	return
}
