package main

import (
	"encoding/json"

	"github.com/boltdb/bolt"

	"io/ioutil"
)

func main() {

	db, _ := bolt.Open("swapi.db", 0600, nil)
	defer db.Close()
	b,_ := ioutil.ReadFile("./fixtures/films.json")
	_films  :=   [6] map[string]interface{}{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("films"))
		if err != nil {
			return err
		}

		for i:=0;i<6;i++{
			_value,_ := json.Marshal(_films[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})



	b,_ = ioutil.ReadFile("./fixtures/people.json")
	_peoples  :=   [82] map[string]interface{}{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("peoples"))
		if err != nil {
			return err
		}

		for i:=0;i<82;i++{
			_value,_ := json.Marshal(_peoples[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/planets.json")
	_planets  :=   [60] map[string]interface{}{}
	json.Unmarshal(b,&_planets)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("planets"))
		if err != nil {
			return err
		}

		for i:=0;i<60;i++{
			_value,_ := json.Marshal(_planets[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})




	b,_ = ioutil.ReadFile("./fixtures/species.json")
	_species  :=   [37] map[string]interface{}{}
	json.Unmarshal(b,&_species)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("species"))
		if err != nil {
			return err
		}

		for i:=0;i<37;i++{
			_value,_ := json.Marshal(_species[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/starships.json")
	_starships  :=   [36] map[string]interface{}{}
	json.Unmarshal(b,&_starships)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("starships"))
		if err != nil {
			return err
		}

		for i:=0;i<36;i++{
			_value,_ := json.Marshal(_starships[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/transport.json")
	_transports :=   [75] map[string]interface{}{}
	json.Unmarshal(b,&_transports)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("transports"))
		if err != nil {
			return err
		}

		for i:=0;i<75;i++{
			_value,_ := json.Marshal(_transports[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/vehicles.json")
	_vehicles  :=   [39] map[string]interface{}{}
	json.Unmarshal(b,&_vehicles)
	db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("vehicles"))
		if err != nil {
			return err
		}

		for i:=0;i<39;i++{
			_value,_ := json.Marshal(_vehicles[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
		}
		return nil
	})
}
