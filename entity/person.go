package entity

import (
	"BringMeFlyServer/database_mysql"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Person struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	Homeworld int    `json:"homeworld"`
	Created   string `json:"created"`
	Edited    string `json:"edited"`
}

func init() {
	personType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"height": &graphql.Field{
				Type: graphql.String,
			},
			"mass": &graphql.Field{
				Type: graphql.String,
			},
			"hair_color": &graphql.Field{
				Type: graphql.String,
			},
			"skin_color": &graphql.Field{
				Type: graphql.String,
			},
			"eye_color": &graphql.Field{
				Type: graphql.String,
			},
			"birth_year": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"homeworld": &graphql.Field{
				Type: graphql.Int,
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

func GetPerson(id int) Person {

	var person Person

	jsonPerson := database_mysql.Search_by_kind_and_id(db, "peoples", id)

	err := json.Unmarshal([]byte(database_mysql.Unmarshal_fields(jsonPerson)), &person)

	if err != nil {
		fmt.Println("error:", err)
	}

	return person
}
