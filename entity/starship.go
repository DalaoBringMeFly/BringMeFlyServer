package entity

import (
	"BringMeFlyServer/database_mysql"
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

func init() {
	starshipType = graphql.NewObject(graphql.ObjectConfig{
		Name: "StarShip",
		Fields: graphql.Fields{
			"hyperdrive_rating": &graphql.Field{
				Type: graphql.String,
			},
			"MGLT": &graphql.Field{
				Type: graphql.String,
			},
			"starship_class": &graphql.Field{
				Type: graphql.String,
			},
			"pilots": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	})
}

func GetStarship(id int) Starship {

	var starship Starship

	jsonStarship := database_mysql.Search_by_kind_and_id(db, "starships", id)

	err := json.Unmarshal([]byte(database_mysql.Unmarshal_fields(jsonStarship)), &starship)
	if err != nil {
		fmt.Println("error:", err)
	}

	return starship
}
