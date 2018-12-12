package entity

import (
	"errors"

	"github.com/boltdb/bolt"
	"github.com/graphql-go/graphql"
)

var db *bolt.DB

func InitDataBase(mdb *bolt.DB) {
	db = mdb
}

func init() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"film": &graphql.Field{
				Type: filmType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return GetFilm(idQuery), nil
					}
					err := errors.New("Field 'film' is missing required arguments: id. ")
					return nil, err
				},
			},
			"person": &graphql.Field{
				Type: personType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return GetPerson(idQuery), nil
					}
					err := errors.New("Field 'person' is missing required arguments: id. ")
					return nil, err
				},
			},
			"planet": &graphql.Field{
				Type: planetType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return GetPlanet(idQuery), nil
					}
					err := errors.New("Field 'planet' is missing required arguments: id. ")
					return nil, err
				},
			},
			"species": &graphql.Field{
				Type: speciesType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return GetSpecies(idQuery), nil
					}
					err := errors.New("Field 'species' is missing required arguments: id. ")
					return nil, err
				},
			},
			"starship": &graphql.Field{
				Type: starshipType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return GetStarship(idQuery), nil
					}
					err := errors.New("Field 'starship' is missing required arguments: id. ")
					return nil, err
				},
			},
			"vehicle": &graphql.Field{
				Type: vehicleType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return GetVehicle(idQuery), nil
					}
					err := errors.New("Field 'vehicle' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	})
}
