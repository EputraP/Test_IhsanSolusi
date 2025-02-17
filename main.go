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

	testService := service.NewTestService(service.TestServiceConfig{
		TestRepo: testRepo,
	})

	testHandler := handler.NewTestHandler(handler.TestHandlerConfig{
		TestService: testService,
	})

	handlers = routes.Handlers{
		TestHandler: testHandler,
	}
	return
}
