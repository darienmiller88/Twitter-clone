package tweed

import (
	"github.com/kamva/mgm/v3"
)

//Tweed is an object meant to define the skeleton of the "tweed" form on the front end.
type Tweed struct{
	mgm.DefaultModel `bson:",inline"`
	Name    string `json:"name" bson:"name"`
	Content string `json:"content" bson:"content"`
}