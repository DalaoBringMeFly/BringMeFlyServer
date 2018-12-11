package entity

import "github.com/graphql-go/graphql"

type Starship struct {
	HyperdriveRating string `json:"hyperdrive_rating"`
	MGLT             string `json:"MGLT"`
	StarshipClass    string `json:"starship_class"`
	PilotURLs        []int  `json:"pilots"`
}

var (
	starshipType      *graphql.Object
	starshipQueryType *graphql.Object
	StarshipSchema    graphql.Schema
)

func init() {

}
