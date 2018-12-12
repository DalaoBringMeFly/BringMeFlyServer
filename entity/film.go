package entity

import (
	"BringMeFlyServer/database"
	"encoding/json"
	"fmt"

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
}

func GetFilm(id int) Film {

	var film Film

	jsonFilm := database.Search_by_kind_and_id(db, "films", id)

	err := json.Unmarshal([]byte(database.Unmarshal_fields(jsonFilm)), &film)
	if err != nil {
		fmt.Println("error:", err)
	}

	return film
}
