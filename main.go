package main

import (
	"fmt"
	"log"
	"os"
	"smart-kost-backend/handler"
	"smart-kost-backend/repository"
	"smart-kost-backend/routes"
	"smart-kost-backend/service"
	dbstore "smart-kost-backend/store/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env")
		log.Fatalln("Error loading env")
	}

	handlers := prepare()

	srv := gin.Default()

	routes.Build(srv, handlers)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := srv.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("Error running gin server: ", err)
		log.Fatalln("Error running gin server: ", err)
	}
}

func prepare() (handlers routes.Handlers) {

	db := dbstore.Get()

	humTempRawRepo := repository.NewHumTempRawRepository(db)

	testService := service.NewTestService()
	humTempRawService := service.NewHumTempRawService(service.HumTempRawServiceConfig{
		HumTempRawRepo: humTempRawRepo,
	})

	test := handler.NewTestHandler(handler.TestHandlerConfig{
		TestService: testService,
	})
	humTempRawHandler := handler.NewHumTempRawHandler(handler.HumTempRawHandlerConfig{
		HumTempRawService: humTempRawService,
	})

	handlers = routes.Handlers{
		Test:       test,
		HumTempRaw: humTempRawHandler,
	}
	return
}
