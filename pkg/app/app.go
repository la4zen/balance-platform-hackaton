package app

import (
	"github.com/golang-jwt/jwt"
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/la4zen/balance-platform-hackaton/pkg/service"
	"github.com/la4zen/balance-platform-hackaton/pkg/store"
	"github.com/la4zen/balance-platform-hackaton/pkg/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	store := store.NewStore()
	service := service.NewService(store)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Static("/img", "static")
	e.POST("/login", service.Login)
	e.PUT("/register", service.Register)
	r := e.Group("/api", middleware.JWT(utils.JWT_KEY), func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("user", &models.User{ID: c.Get("user").(*jwt.Token).Claims.(*models.Claims).ID})
			return next(c)
		}
	})
	r.GET("/accesible", service.Accesible)
	e.Logger.Fatal(e.Start(":8080"))
}
