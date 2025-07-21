package main

import (
	"context"
	"fmt"
	"mp-service/internal/app"
	"mp-service/internal/config"
	"mp-service/internal/database"
	"mp-service/internal/handlers"
	"mp-service/internal/repository/ad"
	"mp-service/internal/repository/user"
	"mp-service/internal/service"
	auth "mp-service/internal/service"
)

func main() {

	cfg, err := config.Load()
	dbUrl := fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", cfg.DB.DB_USER, cfg.DB.DB_PASSWORD, cfg.DB.DB_NAME)
	if err != nil {
		fmt.Println(err.Error())
	}
	db := database.NewDB(dbUrl)
	defer db.DB_CONN.Close(context.Background())

	app := app.NewApp(db)

	userRepo := user.NewUserRepository(db.DB_CONN)
	authService := auth.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	adRepo := ad.NewAdRepository(db.DB_CONN)
	adService := service.NewAdService(adRepo)
	adsHandler := handlers.NewAdsHandler(adService)
	app.MustStart(authHandler, adsHandler)
	//TODO добавить gracefull shutdown
}
