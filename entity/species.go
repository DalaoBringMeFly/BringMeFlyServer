package entity

import "github.com/graphql-go/graphql"

type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	PeopleURLs      []string `json:"people"`
	FilmURLs        []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             string   `json:"url"`
}

var (
	speciesType      *graphql.Object
	speciesQueryType *graphql.Object
	SpeciesSchema    graphql.Schema
)
