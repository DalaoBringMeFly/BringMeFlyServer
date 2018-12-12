package entity

import (
	"BringMeFlyServer/database"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Planet struct {
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surface_water"`
	Population     string `json:"population"`
	Created        string `json:"created"`
	Edited         string `json:"edited"`
}

var (
	planetType      *graphql.Object
	planetQueryType *graphql.Object
	PlanetSchema    graphql.Schema
)

func init() {

}

func GetPlanet(id int) Planet {

	var planet Planet

	jsonPlanet := database.Search_by_kind_and_id(db, "planets", id)

	err := json.Unmarshal([]byte(database.Unmarshal_fields(jsonPlanet)), &planet)
	if err != nil {
		fmt.Println("error:", err)
	}

	return planet
}
