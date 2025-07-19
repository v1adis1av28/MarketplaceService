package main

import (
	"context"
	"fmt"
	"mp-service/internal/app"
	"mp-service/internal/auth"
	"mp-service/internal/config"
	"mp-service/internal/database"
	"mp-service/internal/handlers"
	"mp-service/internal/repository/user"
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
	app.MustStart(authHandler)
	//TODO добавить gracefull shutdown
}
