package routes

import (
	"github.com/gofiber/fiber/v2"
	"webserser/database"
	"webserser/domain"
	"webserser/handlers"
	"webserser/infra"
	"webserser/middleware"
	"webserser/services"
	"webserser/utils"
)

func SetupRoutes(app *fiber.App) {
	setUpOnboardingRoutes(app)
	setUpAuthRoutes(app)
}

func setUpAuthRoutes(app *fiber.App) {
	db := database.GetDatabase(utils.GetValue("DB_NAME"))
	storage, err := infra.NewPostgresStorage(db)
	if err != nil {
		panic(err.Error())
	}
	authService := domain.NewAuthService(storage, services.BcryptHasher{})

	authHandler := handlers.NewAuthHandler(authService)
	authGroup := app.Group("/api/v1/")
	authGroup.Post("login", authHandler.Login)
	app.Use(middleware.NewAuthMiddleware(authService))
	authGroup.Post("logout", authHandler.Logout)
}

func setUpOnboardingRoutes(app *fiber.App) {
	db := database.GetDatabase(utils.GetValue("DB_NAME"))
	storage, err := infra.NewPostgresStorage(db)
	if err != nil {
		panic(err.Error())
	}
	authGroup := app.Group("/api/v1/")

	//authService := domain.NewAuthService(storage, services.BcryptHasher{})
	onboardingService := domain.NewOnboardingService(storage, services.BcryptHasher{})

	onboardingHandler := handlers.NewOnboardingHandler(onboardingService)
	authGroup.Post("signup", onboardingHandler.SignUp)

	//app.Use(middleware.NewAuthMiddleware(authService))

	authGroup.Delete("delete", onboardingHandler.Delete)
}
