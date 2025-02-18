package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EputraP/Test_IhsanSolusi/internal/handler"
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
	"github.com/EputraP/Test_IhsanSolusi/internal/routes"
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
	dbstore "github.com/EputraP/Test_IhsanSolusi/internal/store/db"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
	"github.com/gofiber/fiber"
	"github.com/lpernett/godotenv"
	"github.com/spf13/cobra"
)

var port string
var host string

func main() {
	var rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the API server",
		Run:   startServer,
	}

	rootCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Server host")
	rootCmd.Flags().StringVarP(&port, "port", "P", "8080", "Server port")

	if err := rootCmd.Execute(); err != nil {
		logger.Error("error cobra", "error", err)
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func startServer(cmd *cobra.Command, args []string) {
	if err := logger.Init("app.log"); err != nil {
		fmt.Println("Failed to initialize logger:", err)
		return
	}

	err := godotenv.Load()
	if err != nil {
		logger.Error("error loading env", "error", err)
		log.Fatalln("error loading env", err)
	}

	handlers := prepare()

	srv := fiber.New()

	routes.Build(srv, handlers)

	address := fmt.Sprintf("%s:%s", host, port)
	if err := srv.Listen(address); err != nil {
		logger.Error("Error running fiber server: ", "error", err)
		log.Fatalln("Error running fiber server: ", err)
	}

	logger.Info("Server running", "host", host, "port", port)
}

func prepare() (handlers routes.Handlers) {
	db := dbstore.Get()

	userRepo := repository.NewUserRepository(db)
	userSaldoRepo := repository.NewUserSaldoRepository(db)

	userService := service.NewUserService(service.UserServiceConfig{
		UserRepo:      userRepo,
		UserSaldoRepo: userSaldoRepo,
	})
	userSaldoService := service.NewUserSaldoService(service.UserSaldoServiceConfig{
		UserSaldoRepo: userSaldoRepo,
		UserRepo:      userRepo,
	})

	userHandler := handler.NewUserHandler(handler.UserHandlerConfig{
		UserService: userService,
	})
	userSaldoHandler := handler.NewUserSaldoHandler(handler.UserSaldoHandlerConfig{
		UserSaldoService: userSaldoService,
	})

	handlers = routes.Handlers{
		UserHandler:      userHandler,
		UserSaldoHandler: userSaldoHandler,
	}
	return
}
