package db
import(
	"fmt"
	"github.com/boltdb/bolt"
	"strings"
	"encoding/json"
)

func Search_by_kind_and_id(db *bolt.DB,kind string,id int)string{
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(kind))
		v = b.Get([]byte(string(id)))
		//fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	return string(v)
}

func Search_by_kind_and_keyword(db *bolt.DB,kind string,keyword string)[]string{
	var string_array []string
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(kind))
		b.ForEach(func(k, v []byte) error {
			if(strings.Contains(string(v),keyword)==true){
				string_array = append(string_array,string(v))
			}
			//fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	return string_array
}

func Search_by_keyword(db *bolt.DB,keyword string)[]string{
	var string_array []string
	var kinds = []string{"films","peoples","planets","species","starships","transports","vehicles"}
	for i:=0;i<7;i++{
	
		db.View(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte(kinds[i]))
			b.ForEach(func(k, v []byte) error {
				if(strings.Contains(string(v),keyword)==true){
					string_array = append(string_array,string(v))
				}
				//fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			return nil
		})
	}
	return string_array
}


func Example(){
	db,_ := bolt.Open("swapi.db",0600,nil)
	defer db.Close()

	s := Search_by_keyword(db ,"Skywalker")
	fmt.Println(s)
}

func Unmarshal_fields(s string)string{
	var _json map[string]interface{}
	json.Unmarshal([]byte(s),&_json)
	//fmt.Println(_json)
	temp,_ := json.Marshal(_json["fields"])
	//fmt.Println(temp)
	return string(temp)
	//return value
}

func test(){
	db,_ := bolt.Open("swapi.db",0600,nil)
	s := Search_by_kind_and_id(db,"films",1)
	//fmt.Println(s)
	fmt.Println(Unmarshal_fields(string(s)))
}