package app

import (
	"log"
	"mp-service/internal/database"
	"mp-service/internal/handlers"
	"mp-service/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *database.DB
	Router *gin.Engine
}

func NewApp(db *database.DB) *App {
	return &App{
		DB:     db,
		Router: gin.Default(),
	}
}

func (a *App) MustStart(ah *handlers.AuthHandler, ads *handlers.AdsHandler) {
	a.Router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	if err := a.Run(ah, ads); err != nil {
		panic(err)
	}
}

func (app *App) Run(ah *handlers.AuthHandler, ads *handlers.AdsHandler) error {

	if err := app.SetupRoutes(ah, ads); err != nil {
		log.Fatal("Failed to setup server routes", "error", err)
		return err
	}

	if err := app.Router.Run(); err != nil {
		log.Fatal("Failed to start server", "error", err)
		return err
	}
	return nil
}

func (app *App) SetupRoutes(ah *handlers.AuthHandler, ads *handlers.AdsHandler) error {
	app.Router.POST("/auth/login", ah.Login)
	app.Router.POST("/auth/register", ah.Register)
	app.Router.GET("/auth/logout", ah.Logout)

	app.Router.POST("/api/advertisement", middleware.AuthMiddleware(), ads.CreateAd)
	app.Router.GET("/api/advertisements", ads.GetAds)

	return nil
}
