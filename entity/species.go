package entity

import (
	"BringMeFlyServer/database"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Species struct {
	Name            string `json:"name"`
	Classification  string `json:"classification"`
	Designation     string `json:"designation"`
	AverageHeight   string `json:"average_height"`
	SkinColors      string `json:"skin_colors"`
	HairColors      string `json:"hair_colors"`
	EyeColors       string `json:"eye_colors"`
	AverageLifespan string `json:"average_lifespan"`
	Homeworld       string `json:"homeworld"`
	Language        string `json:"language"`
	PeopleURLs      []int  `json:"people"`
	Created         string `json:"created"`
	Edited          string `json:"edited"`
}

var (
	speciesType      *graphql.Object
	speciesQueryType *graphql.Object
	SpeciesSchema    graphql.Schema
)

func init() {

}

func GetSpecies(id int) Species {

	var species Species

	jsonSpecies := database.Search_by_kind_and_id(db, "species", id)

	err := json.Unmarshal([]byte(database.Unmarshal_fields(jsonSpecies)), &species)
	if err != nil {
		fmt.Println("error:", err)
	}

	return species
}
