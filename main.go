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
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
	"github.com/gofiber/fiber"
	"github.com/lpernett/godotenv"
)

func main() {

	// Initialize the global logger
	if err := logger.Init("app.log"); err != nil {
		fmt.Println("Failed to initialize logger:", err)
		return
	}

	env := os.Getenv(constant.EnvKeyEnv)

	if env != "prod" {
		err := godotenv.Load()

		if err != nil {
			logger.Error("error loading env", "error", err)
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
		logger.Error("Error running gin server: ", "error", err)
		log.Fatalln("Error running gin server: ", err)
	}
	logger.Info("Server running", "port", port)
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
		UserRepo:      userRepo,
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
