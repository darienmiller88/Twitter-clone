package tweedcontroller

import (
	"fmt"
	"net/http"
	"twitter_clone/models"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_echo"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//TweedController will handle routes that will accept incoming tweeds from the front end
type TweedController struct{

}

//Init - This exported method will initialize each route with the routeGroup parameter, mounting each route on 
//the "/api" route set in main.
func (t *TweedController) Init(routeGroup *echo.Group){
	t.addTweed(routeGroup)
	t.getTweeds(routeGroup)
	t.getTweedsByName(routeGroup)
	t.deleteTweeds(routeGroup)
	t.updateTweed(routeGroup)
}

//AddPost will add a tweed to the database.
func (t *TweedController) addTweed(routeGroup *echo.Group){
	//Rate limiter to limit each request to one per second.
	limiter := tollbooth.NewLimiter(1, nil)
	limiter.SetMessage("limit reached")

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
	}, tollbooth_echo.LimitHandler(limiter))
}

func (t *TweedController) getTweedsByName(routeGroup *echo.Group){
	routeGroup.Add("GET", "/tweeds/:name", func(c echo.Context) error {	
		//create a slice of Tweed objects to capture each Tweed from the mongo database
		docsArray := []models.Tweed{}

		//Retrieve the documents from colletions using the query parameter
		documents, _ := mgm.CollectionByName("tweeds").Find(mgm.Ctx(), bson.M{"name": c.Param("name")})

		//Unmarshal the retrieved documents in the array!
		documents.All(mgm.Ctx(), &docsArray)
		return c.JSONPretty(http.StatusOK, docsArray, " ")
	})
}

func (t *TweedController) getTweeds(routeGroup *echo.Group){
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

func (t *TweedController) deleteTweeds(routeGroup *echo.Group){
	//For security, I chose to comment on the delete route upon deployment, saving it for local usage to
	// test CRUD functionality.

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
	// 		"Tweed deleted!"    : userTweed,
	// 		"Num tweeds deleted": result.DeletedCount,
	// 	}, " ")
	// })
}

func (t *TweedController) updateTweed(routeGroup *echo.Group){
	//As with the delete route, for deployment, I chose to comment out the update, but it can certaintly
	//be uncommented for local usage

	// routeGroup.Add("PUT", "/tweeds", func(c echo.Context) error {	
	// 	updateParameters := make(map[string]string)

	// 	errHandle(c.Bind(&updateParameters))
	// 	if updateParameters["name"] == "" || updateParameters["content"] == ""{
	// 		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
	// 			"error": "Your tweed MUST have a name AND body!",
	// 			"tweed sent": echo.Map{
	// 				"name":    updateParameters["name"],
	// 				"content": updateParameters["content"],
	// 			},
	// 		})
	// 	}

	//  	result, err := mgm.CollectionByName("tweeds").UpdateOne(
	// 		//Send in a default context for a time out
	// 		mgm.Ctx(),

	// 		//These are the filter parameters we will query our database for. Find a tweed with a name and content
	// 		//that matches the one sent to this route.
	// 		bson.M{"name": updateParameters["name"], "content": updateParameters["content"]},

	// 		//Finally, apply the following parameters to the found tweed! 
	// 		bson.M{ "$set": bson.M{
	// 					"name": updateParameters["updated_name"], "content": updateParameters["updated_content"],
	// 				},
	// 		},
	// 	)

	// 	fmt.Println(err)

	// 	//Return the result of the query.
	// 	return c.JSONPretty(http.StatusOK, result, " ")
	// })
}

//Function to validate an incoming tweed from the front end. Checks to see if both the name and content bodies are
//filled out.
func validateTweed(userTweed models.Tweed) bool{
	return len(userTweed.Name) != 0 && len(userTweed.Content) != 0
}

func errHandle(err error){
	if err != nil{
		fmt.Println(err)
	}
}