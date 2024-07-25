package main

import (
	"fmt"
	"log"
	"os"
	"smart-kost-backend/constant"
	"smart-kost-backend/handler"
	"smart-kost-backend/middleware"
	"smart-kost-backend/repository"
	"smart-kost-backend/routes"
	"smart-kost-backend/service"
	dbstore "smart-kost-backend/store/db"
	"smart-kost-backend/util/hasher"
	"smart-kost-backend/util/tokenprovider"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env")
		log.Fatalln("Error loading env")
	}

	handlers, middlewares := prepare()

	srv := gin.Default()

	srv.Use(middleware.CORS())

	routes.Build(srv, handlers, middlewares)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := srv.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("Error running gin server: ", err)
		log.Fatalln("Error running gin server: ", err)
	}
}

func prepare() (handlers routes.Handlers, middlewares routes.Middlewares) {

	appName := os.Getenv(constant.EnvKeyAppName)
	jwtSecret := os.Getenv(constant.EnvKeyJWTSecret)
	refreshTokenDurationStr := os.Getenv(constant.EnvKeyRefreshTokenDuration)

	accessTokenDurationStr := os.Getenv(constant.EnvKeyAccessTokenDuration)

	refreshTokenDuration, err := strconv.Atoi(refreshTokenDurationStr)

	if err != nil {
		log.Fatalln("error creating handlers and middleware", err)
	}

	accessTokenDuration, err := strconv.Atoi(accessTokenDurationStr)
	if err != nil {
		log.Fatalln("error creating handlers and middlewares", err)
	}

	jwtProvider := tokenprovider.NewJWT(appName, jwtSecret, refreshTokenDuration, accessTokenDuration)

	middlewares = routes.Middlewares{
		Auth: middleware.CreateAuth(jwtProvider),
	}

	db := dbstore.Get()

	humTempRawRepo := repository.NewHumTempRawRepository(db)
	userRepo := repository.NewUserRepository(db)

	testService := service.NewTestService()
	humTempRawService := service.NewHumTempRawService(service.HumTempRawServiceConfig{
		HumTempRawRepo: humTempRawRepo,
	})
	hasher := hasher.NewBcrypt(10)
	authService := service.NewUserService(service.AuthServiceConfig{
		UserRepo:    userRepo,
		Hasher:      hasher,
		JwtProvider: tokenprovider.GetProvider(),
	})

	test := handler.NewTestHandler(handler.TestHandlerConfig{
		TestService: testService,
	})
	humTempRawHandler := handler.NewHumTempRawHandler(handler.HumTempRawHandlerConfig{
		HumTempRawService: humTempRawService,
	})
	authHandler := handler.NewAuthRawHandler(handler.AuthHandlerConfig{
		AuthService: authService,
	})

	handlers = routes.Handlers{
		Test:       test,
		HumTempRaw: humTempRawHandler,
		Auth:       authHandler,
	}
	return
}
