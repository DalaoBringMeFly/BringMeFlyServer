package entity

import (
	"BringMeFlyServer/database_mysql"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Species struct {
	Name            string `json:"name"`
	Classification  string `json:"classification"`
	Designation     string `json:"designation"`
	AverageHeight   string `json:"average_height"`
	SkinColors      string `json:"skin_colors"`
	HairColors      string `json:"hair_colors"`
	EyeColors       string `json:"eye_colors"`
	AverageLifespan string `json:"average_lifespan"`
	Homeworld       int    `json:"homeworld"`
	Language        string `json:"language"`
	PeopleURLs      []int  `json:"people"`
	Created         string `json:"created"`
	Edited          string `json:"edited"`
}

func init() {
	speciesType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Species",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"classification": &graphql.Field{
				Type: graphql.String,
			},
			"designation": &graphql.Field{
				Type: graphql.String,
			},
			"average_height": &graphql.Field{
				Type: graphql.String,
			},
			"skin_colors": &graphql.Field{
				Type: graphql.String,
			},
			"hair_colors": &graphql.Field{
				Type: graphql.String,
			},
			"eye_colors": &graphql.Field{
				Type: graphql.String,
			},
			"average_lifespan": &graphql.Field{
				Type: graphql.String,
			},
			"homeworld": &graphql.Field{
				Type: graphql.Int,
			},
			"language": &graphql.Field{
				Type: graphql.String,
			},
			"people": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
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

func GetSpecies(id int) Species {

	var species Species

	jsonSpecies := database_mysql.Search_by_kind_and_id(db, "species", id)

	err := json.Unmarshal([]byte(database_mysql.Unmarshal_fields(jsonSpecies)), &species)
	if err != nil {
		fmt.Println("error:", err)
	}

	return species
}
