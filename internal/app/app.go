package app

import (
	"log"
	"os"
	"workshop-storage-api-docs/internal/controller/rest"
	"workshop-storage-api-docs/internal/repository"
	"workshop-storage-api-docs/internal/usecase"
	"workshop-storage-api-docs/pkg/bcrypt"
	httpserver "workshop-storage-api-docs/pkg/gin"
	"workshop-storage-api-docs/pkg/jwt"
	"workshop-storage-api-docs/pkg/middleware"
	"workshop-storage-api-docs/pkg/oauth"
	"workshop-storage-api-docs/pkg/postgres"
	"workshop-storage-api-docs/pkg/supabase"

	"github.com/go-playground/validator/v10"
)

func Run() {
	db := postgres.StartPostgres()
	app := httpserver.Start()
	jwtInit := *jwt.NewJWT()
	bcryptInit := bcrypt.NewBcrypt()
	storageClient := supabase.NewStorageClient()
	middleware := middleware.NewMiddleware(&jwtInit)
	validator := validator.New()

	repo := repository.NewRepository(db)
	oauthConfig := oauth.GoogleOAuthConfig()
	uc := usecase.NewUsecase(jwtInit, bcryptInit, &oauthConfig, storageClient, repo)
	v1 := rest.NewV1(middleware, validator, uc)

	rest.NewRouter(app, v1)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
