package main

import (
	"fmt"
	"os"
	"time"
	"twitter_clone/middlewares"
	"twitter_clone/defaultTemplate"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	errHandle(godotenv.Load())
	mgm.SetDefaultConfig(nil, "", options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
	app := echo.New()

	app.Static("/public", "public")
	mymiddlewares.AttachMiddleWares(app)
	app.Renderer = getdefaulttemplate.GetRenderer("templates/*.html")

	app.GET("/:name", func(c echo.Context) error {
		fmt.Println(c.Param("name"))
		return c.Render(200, "index.html", echo.Map{
			"Name": c.Param("name"),
			"Time": time.Now().Format("3:4 PM"),
		})
	})

	fmt.Println("Server running")
	app.Logger.Fatal(app.Start(":8080"))
}

func errHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
