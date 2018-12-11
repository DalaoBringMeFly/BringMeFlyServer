package entity

import "github.com/graphql-go/graphql"

type Vehicle struct {
	VehicleClass string `json:"vehicle_class"`
	PilotURLs    []int  `json:"pilots"`
}

var (
	vehicleType      *graphql.Object
	vehicleQueryType *graphql.Object
	VehicleSchema    graphql.Schema
)
