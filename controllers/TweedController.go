package tweedcontroller

import (
	"fmt"
	"net/http"
	"twitter_clone/models"

	"github.com/labstack/echo/v4"
)

//TweedController will handle routes that will accept incoming tweeds from the front end
type TweedController struct{

}

//Init - This exported method will initialize each route.
func (p *TweedController) Init(routeGroup *echo.Group){
	p.addPost(routeGroup)
	p.getPost(routeGroup)
}

//AddPost will add a tweed to the database.
func (p *TweedController) addPost(routeGroup *echo.Group){
	routeGroup.Add("POST", "/tweed", func(c echo.Context) error {
		userTweed := tweed.Tweed{}
		errHandle(c.Bind(&userTweed))

		if validateTweed(userTweed){
			fmt.Println(userTweed)
			return c.JSON(http.StatusCreated, userTweed)
		}
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "Your tweed MUST have a name and body!",
		})
	})
}

func (p *TweedController) getPost(routeGroup *echo.Group){
	routeGroup.Add("GET", "/tweed", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello from api/tweed",
		})
	})
}

//Function to validate an incoming tweed from the front end. The map will be guaranteed to contain to key - value
//pairs: A "name" key with a string name, and a "content" key with a string containing the body of the tweed.
func validateTweed(userTweed tweed.Tweed) bool{
	return len(userTweed.Name) != 0 && len(userTweed.Content) != 0
}

func errHandle(err error){
	if err != nil{
		fmt.Println(err)
	}
}