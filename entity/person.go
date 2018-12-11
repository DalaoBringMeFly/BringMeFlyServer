package entity

import (
	"errors"

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
	Homeworld string `json:"homeworld"`
	Created   string `json:"created"`
	Edited    string `json:"edited"`
}

var (
	personType      *graphql.Object
	personQueryType *graphql.Object
	PersonSchema    graphql.Schema
)

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
				Type: graphql.String,
			},
			"films": &graphql.Field{
				Type: graphql.String,
			},
			"species": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"starships": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"created": &graphql.Field{
				Type: graphql.String,
			},
			"edited": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	personQueryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"persons": &graphql.Field{
				Type: personType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					/*idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return GetPerson(p.Args["id"].(int)), nil
					}*/
					err := errors.New("Field 'persons' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	})

	PersonSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: personQueryType,
	})
}

func GetPerson(id int) Person {
	return Person{}
}
