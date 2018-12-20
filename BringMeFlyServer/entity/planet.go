package entity

import (
	"BringMeFlyServer/database"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Planet struct {
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surface_water"`
	Population     string `json:"population"`
	Created        string `json:"created"`
	Edited         string `json:"edited"`
}

func init() {
	planetType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Planet",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"rotation_period": &graphql.Field{
				Type: graphql.String,
			},
			"orbital_period": &graphql.Field{
				Type: graphql.String,
			},
			"diameter": &graphql.Field{
				Type: graphql.String,
			},
			"climate": &graphql.Field{
				Type: graphql.String,
			},
			"gravity": &graphql.Field{
				Type: graphql.String,
			},
			"terrain": &graphql.Field{
				Type: graphql.String,
			},
			"surface_water": &graphql.Field{
				Type: graphql.String,
			},
			"population": &graphql.Field{
				Type: graphql.String,
			},
			"created": &graphql.Field{
				Type: graphql.String,
			},
			"edited": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func GetPlanet(id int) Planet {

	var planet Planet

	jsonPlanet := database.Search_by_kind_and_id(db, "planets", id)

	err := json.Unmarshal([]byte(database.Unmarshal_fields(jsonPlanet)), &planet)
	if err != nil {
		fmt.Println("error:", err)
	}

	return planet
}
