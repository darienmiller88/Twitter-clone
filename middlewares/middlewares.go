package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//AttachMiddleWares is a function that will attached the necessary middlewares for the router.
func AttachMiddleWares(app *echo.Echo){
	app.Use(middleware.Recover())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))
	app.Use(middleware.CORS())
}