package main

import (
	"os"
	//"twitter_clone/controllers"
	"twitter_clone/defaultTemplate"
	"twitter_clone/middlewares"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {	
	//First, load the env file
	godotenv.Load()	
	
	//Using the mgm package, load the default config settings for the mongodb client conneciton.
	mgm.SetDefaultConfig(nil, os.Getenv("DB_NAME"), options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))

	//Declare both the default router and a route group for each route to be mounted on ("/api")
	app := echo.New()
	//api := app.Group("/api")

	//Initialize the "tweed" controller, which will handle the posting and retrieving of tweeds
	//tweedRoutes := tweedcontroller.TweedController{}
	app.Static("/public", "public")

	//Echo does not use the default golang template which is my preference, so this line of code will assign the 
	//Echo renderer an instance of that template engine
	app.Renderer = getdefaulttemplate.GetRenderer("templates/*.html")

	//Call my middlewares package to attach global middleware to the Echo app
	middlewares.AttachMiddleWares(app)

	//Mount each route of the "tweedcontroller" onto the /api route.
	//tweedRoutes.Init(api)

	app.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", echo.Map{})
	})

	app.Logger.Fatal(app.Start(":1323"))
}