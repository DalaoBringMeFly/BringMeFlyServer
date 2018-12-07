package entity

import (
	"errors"

	"github.com/graphql-go/graphql"
)

type Film struct {
	Title         string   `json:"title"`
	EpisodeID     int      `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	StarshipURLs  []string `json:"starships"`
	VehicleURLs   []string `json:"vehicles"`
	SpeciesURLs   []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	URL           string   `json:"url"`
}

var (
	filmType      *graphql.Object
	filmQueryType *graphql.Object
	FilmSchema    graphql.Schema
)

func init() {
	filmType = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Film",
		Fields: graphql.Fields{},
	})

	filmQueryType = graphql.NewObject(graphql.ObjectConfig{
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
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return GetPerson(p.Args["id"].(int)), nil
					}
					err := errors.New("Field 'persons' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	})

	FilmSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: filmQueryType,
	})
}

func GetPerson(id int) Person {
	if person, ok := PersonData[id]; ok {
		return person
	}
	return Person{}
}
