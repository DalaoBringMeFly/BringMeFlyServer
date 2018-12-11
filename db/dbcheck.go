package main

import(
	//"fmt"
	"github.com/boltdb/bolt"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
)
func main(){
	//fmt.Println(_films[0]["model"])
	//fmt.Printf("yest")
	

	db,_ := bolt.Open("swapi.db",0600,nil)
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("planets"))
		v := b.Get([]byte(string(1)))
		//fmt.Printf("The answer is: %s\n", v)
		fmt.Println(string(v))
		return nil
	})

}

