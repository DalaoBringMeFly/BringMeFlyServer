package database_mysql

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"io/ioutil"
)

func main() {
	db := OpenDatabase("swapi")
	defer db.Close()
	//films
	CreateTable(db, "films")
	b, _ := ioutil.ReadFile("../database/fixtures/films.json")
	_films := [6]map[string]interface{}{}
	json.Unmarshal(b, &_films)

	tx, _ := db.Begin()
	for i := 0; i < 6; i++ {
		_value, _ := json.Marshal(_films[i])
		tx.Exec("INSERT INTO films(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	//peoples
	CreateTable(db, "peoples")
	b, _ = ioutil.ReadFile("../database/fixtures/people.json")
	_peoples := [82]map[string]interface{}{}
	json.Unmarshal(b, &_peoples)

	tx, _ = db.Begin()
	for i := 0; i < 82; i++ {
		_value, _ := json.Marshal(_peoples[i])
		tx.Exec("INSERT INTO peoples(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	//planets
	CreateTable(db, "planets")
	b, _ = ioutil.ReadFile("../database/fixtures/planets.json")
	_planets := [60]map[string]interface{}{}
	json.Unmarshal(b, &_planets)

	tx, _ = db.Begin()
	for i := 0; i < 60; i++ {
		_value, _ := json.Marshal(_planets[i])
		tx.Exec("INSERT INTO planets(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	//species
	CreateTable(db, "species")
	b, _ = ioutil.ReadFile("../database/fixtures/species.json")
	_species := [37]map[string]interface{}{}
	json.Unmarshal(b, &_species)

	tx, _ = db.Begin()
	for i := 0; i < 37; i++ {
		_value, _ := json.Marshal(_species[i])
		tx.Exec("INSERT INTO species(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	//starships
	CreateTable(db, "starships")
	b, _ = ioutil.ReadFile("../database/fixtures/starships.json")
	_starships := [36]map[string]interface{}{}
	json.Unmarshal(b, &_starships)

	tx, _ = db.Begin()
	for i := 0; i < 36; i++ {
		_value, _ := json.Marshal(_starships[i])
		tx.Exec("INSERT INTO starships(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	//transports
	CreateTable(db, "transports")
	b, _ = ioutil.ReadFile("../database/fixtures/transport.json")
	_transports := [75]map[string]interface{}{}
	json.Unmarshal(b, &_transports)

	tx, _ = db.Begin()
	for i := 0; i < 75; i++ {
		_value, _ := json.Marshal(_transports[i])
		tx.Exec("INSERT INTO transports(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	//vehicles
	CreateTable(db, "vehicles")
	b, _ = ioutil.ReadFile("../database/fixtures/vehicles.json")
	_vehicles := [39]map[string]interface{}{}
	json.Unmarshal(b, &_vehicles)

	tx, _ = db.Begin()
	for i := 0; i < 39; i++ {
		_value, _ := json.Marshal(_vehicles[i])
		tx.Exec("INSERT INTO vehicles(id, data) values(?,?)", i, _value)
	}
	tx.Commit()

	fmt.Printf(Search_by_kind_and_id(db, "films", 2))
}
