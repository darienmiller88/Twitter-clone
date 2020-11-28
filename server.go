package main

import (
	"fmt"
	"os"
	"twitter_clone/defaultTemplate"
	"twitter_clone/middlewares"
	"twitter_clone/controllers"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()
	mgm.SetDefaultConfig(nil, os.Getenv("DB_NAME"), options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
	app := echo.New()
	api := app.Group("/api")
	tweedRoutes := tweedcontroller.TweedController{}

	fmt.Println(os.Getenv("DB_CONNECTION"))
	app.Static("/public", "public")
	app.Renderer = getdefaulttemplate.GetRenderer("templates/*.html")
	middlewares.AttachMiddleWares(app)
	tweedRoutes.Init(api)
	
	app.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", echo.Map{})
	})

	fmt.Println("Server running")
	app.Logger.Fatal(app.Start(":8080"))
}