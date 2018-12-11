package main

import(
	//"fmt"
	"github.com/boltdb/bolt"
	"encoding/json"
	"fmt"
	"io/ioutil"
)
func main(){
	//fmt.Println(_films[0]["model"])
	//fmt.Printf("yest")
	

	db,_ := bolt.Open("_swapi.db",0600,nil)
	defer db.Close()
	b,_ := ioutil.ReadFile("./fixtures/films.json")
	_films  :=   [6] map[string]interface{}{}
	json.Unmarshal(b,&_films)
	db.Update(func(tx *bolt.Tx) error {
		bu,err := tx.CreateBucketIfNotExists([]byte("films"))
		if err!=nil{
			return err
		}

		for i:=0;i<6;i++{
			_value,_ := json.Marshal(_films[i])
			bu.Put([]byte(string(i)),[]byte(string(_value)))
			fmt.Printf(string(_value))
		}
		return nil
	})

}