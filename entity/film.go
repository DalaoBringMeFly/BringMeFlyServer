package entity

import (
	"errors"
	"log"

	"github.com/boltdb/bolt"
	"github.com/graphql-go/graphql"
)

type Film struct {
	Producer      string `json:"producer"`
	Title         string `json:"title"`
	EpisodeID     int    `json:"episode_id"`
	OpeningCrawl  string `json:"opening_crawl"`
	Director      string `json:"director"`
	CharacterURLs []int  `json:"characters"`
	PlanetURLs    []int  `json:"planets"`
	StarshipURLs  []int  `json:"starships"`
	VehicleURLs   []int  `json:"vehicles"`
	SpeciesURLs   []int  `json:"species"`
	Created       string `json:"created"`
	Edited        string `json:"edited"`
	ReleaseData   string `json:"release_date"`
}

var (
	filmType      *graphql.Object
	filmQueryType *graphql.Object
	FilmSchema    graphql.Schema
)

func init() {
	filmType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Film",
		Fields: graphql.Fields{
			"producer": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"episode_id": &graphql.Field{
				Type: graphql.Int,
			},
			"opening_crawl": &graphql.Field{
				Type: graphql.String,
			},
			"director": &graphql.Field{
				Type: graphql.String,
			},
			"characters": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"planets": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"starships": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"species": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"created": &graphql.Field{
				Type: graphql.String,
			},
			"edited": &graphql.Field{
				Type: graphql.String,
			},
			"release_date": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	filmQueryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"film": &graphql.Field{
				Type: personType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return GetFilm(idQuery), nil
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

func GetFilm(id string) Film {

	film := Film{}

	db, err := bolt.Open("swapi.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var f []byte = nil

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("films"))
		f = b.Get([]byte(id))
		return nil
	})

	if f != nil {

	}

	return film
}
