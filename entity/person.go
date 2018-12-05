package entity

import "github.com/graphql-go/graphql"

type Person struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Height       string   `json:"height"`
	Mass         string   `json:"mass"`
	HairColor    string   `json:"hair_color"`
	SkinColor    string   `json:"skin_color"`
	EyeColor     string   `json:"eye_color"`
	BirthYear    string   `json:"birth_year"`
	Gender       string   `json:"gender"`
	Homeworld    string   `json:"homeworld"`
	FilmURLs     []string `json:"films"`
	SpeciesURLs  []string `json:"species"`
	VehicleURLs  []string `json:"vehicles"`
	StarshipURLs []string `json:"starships"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

var personsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
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
	},
)
var personsListType = graphql.NewList(personsType)
