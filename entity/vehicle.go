package entity

import "github.com/graphql-go/graphql"

type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	PilotURLs            []string `json:"pilots"`
	FilmURLs             []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

var (
	vehicleType      *graphql.Object
	vehicleQueryType *graphql.Object
	VehicleSchema    graphql.Schema
)
