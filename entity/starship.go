package entity

import (
	"BringMeFlyServer/database"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Starship struct {
	HyperdriveRating string `json:"hyperdrive_rating"`
	MGLT             string `json:"MGLT"`
	StarshipClass    string `json:"starship_class"`
	PilotURLs        []int  `json:"pilots"`
}

var (
	starshipType      *graphql.Object
	starshipQueryType *graphql.Object
	StarshipSchema    graphql.Schema
)

func init() {

}

func GetStarship(id int) Starship {

	var starship Starship

	jsonStarship := database.Search_by_kind_and_id(db, "starships", id)

	err := json.Unmarshal([]byte(database.Unmarshal_fields(jsonStarship)), &starship)
	if err != nil {
		fmt.Println("error:", err)
	}

	return starship
}
