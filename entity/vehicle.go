package entity

import (
	"BringMeFlyServer/database"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Vehicle struct {
	VehicleClass string `json:"vehicle_class"`
	PilotURLs    []int  `json:"pilots"`
}

var (
	vehicleType      *graphql.Object
	vehicleQueryType *graphql.Object
	VehicleSchema    graphql.Schema
)

func init() {

}

func GetVehicle(id int) Vehicle {

	var vehicle Vehicle

	jsonVehicle := database.Search_by_kind_and_id(db, "vehicles", id)

	err := json.Unmarshal([]byte(database.Unmarshal_fields(jsonVehicle)), &vehicle)
	if err != nil {
		fmt.Println("error:", err)
	}

	return vehicle
}
