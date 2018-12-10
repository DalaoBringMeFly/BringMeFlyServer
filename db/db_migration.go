package main

import(
	//"fmt"
	"github.com/boltdb/bolt"
	"encoding/json"

	"io/ioutil"
)
func main(){
	//fmt.Println(_films[0]["model"])
	//fmt.Printf("yest")
	

	db,_ := bolt.Open("swapi.db",0600,nil)
	defer db.Close()
	b,_ := ioutil.ReadFile("./fixtures/films.json")
	_films  :=   [6] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("films"))
		if err!=nil{
			return err
		}

		for i:=0;i<6;i++{
			bu.Put([]byte(string(i)),[]byte(_films[i]))
		}
		return nil
	})



	b,_ = ioutil.ReadFile("./fixtures/people.json")
	_peoples  :=   [82] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("peoples"))
		if err!=nil{
			return err
		}

		for i:=0;i<82;i++{
			bu.Put([]byte(string(i)),[]byte(_peoples[i]))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/planets.json")
	_planets  :=   [60] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("planets"))
		if err!=nil{
			return err
		}

		for i:=0;i<60;i++{
			bu.Put([]byte(string(i)),[]byte(_planets[i]))
		}
		return nil
	})




	b,_ = ioutil.ReadFile("./fixtures/species.json")
	_species  :=   [37] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("species"))
		if err!=nil{
			return err
		}

		for i:=0;i<37;i++{
			bu.Put([]byte(string(i)),[]byte(_species[i]))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/starships.json")
	_starships  :=   [36] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("starships"))
		if err!=nil{
			return err
		}

		for i:=0;i<36;i++{
			bu.Put([]byte(string(i)),[]byte(_starships[i]))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/transport.json")
	_transports :=   [75] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("transports"))
		if err!=nil{
			return err
		}

		for i:=0;i<75;i++{
			bu.Put([]byte(string(i)),[]byte(_transports[i]))
		}
		return nil
	})

	b,_ = ioutil.ReadFile("./fixtures/vehicles.json")
	_vehicles  :=   [39] string{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("vehicles"))
		if err!=nil{
			return err
		}

		for i:=0;i<39;i++{
			bu.Put([]byte(string(i)),[]byte(_vehicles[i]))
		}
		return nil
	})
}