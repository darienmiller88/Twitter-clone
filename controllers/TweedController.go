package tweedcontroller

import (
	"fmt"
	"net/http"
	"twitter_clone/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//TweedController will handle routes that will accept incoming tweeds from the front end
type TweedController struct{

}

//Init - This exported method will initialize each route.
func (p *TweedController) Init(routeGroup *echo.Group){
	p.addTweed(routeGroup)
	p.getTweeds(routeGroup)
	p.deleteTweed(routeGroup)
}

//AddPost will add a tweed to the database.
func (p *TweedController) addTweed(routeGroup *echo.Group){
	routeGroup.Add("POST", "/tweeds", func(c echo.Context) error {
		userTweed := models.Tweed{}

		//Bind the above Tweed object to the request body, which will be a JSON object from the front end 
		//containing a name and a content body.
		errHandle(c.Bind(&userTweed))

		//If the user's tweed contains both a name AND content, insert the tweed into the database, and respond
		//with the user's tweed along with a status code of 201 (created). Otherwise, respond with a status code
		//of 422 and a message telling the user their tweed is incomplete.
		if validateTweed(userTweed){
			errHandle(mgm.Coll(&userTweed).Create(&userTweed))
			return c.JSON(http.StatusCreated, userTweed)
		}
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "Your tweed MUST have a name and body!",
		})
	})
}

func (p *TweedController) getTweeds(routeGroup *echo.Group){
	routeGroup.Add("GET", "/tweeds", func(c echo.Context) error {	
		//create a slice of Tweed objects to capture each Tweed from the mongo database
		docsArray := []models.Tweed{}

		//Retrieve the documents from colletions
		documents, _ := mgm.CollectionByName("tweeds").Find(mgm.Ctx(), bson.D{})

		//Send the retrieved documents in the array!
		documents.All(mgm.Ctx(), &docsArray)
		return c.JSONPretty(http.StatusOK, docsArray, " ")
	})
}

func (p *TweedController) deleteTweed(routeGroup *echo.Group){
	// routeGroup.Add("DELETE", "/tweeds", func(c echo.Context) error {	
	// 	userTweed := models.Tweed{}

	// 	errHandle(c.Bind(&userTweed))

	// 	if userTweed.Name == ""{
	// 		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
	// 			"error": "Your tweed MUST have a name!",
	// 		})
	// 	}

		
	// 	result, _ := mgm.CollectionByName("tweeds").DeleteMany(mgm.Ctx(), bson.M{"name": userTweed.Name})
	// 	return c.JSONPretty(http.StatusOK, echo.Map{
	// 		"Tweed deleted!": userTweed,
	// 		"Num tweeds deleted": result.DeletedCount,
	// 	}, " ")
	// })
}

//Function to validate an incoming tweed from the front end. The map will be guaranteed to contain to key - value
//pairs: A "name" key with a string name, and a "content" key with a string containing the body of the tweed.
func validateTweed(userTweed models.Tweed) bool{
	return len(userTweed.Name) != 0 && len(userTweed.Content) != 0
}

func errHandle(err error){
	if err != nil{
		fmt.Println(err)
	}
}