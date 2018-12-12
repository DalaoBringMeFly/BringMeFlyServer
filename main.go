package main

import (
	"BringMeFlyServer/entity"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/graphql-go/graphql"
)

func main() {
	db, err := bolt.Open("./database/swapi.db", 0600, nil)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer db.Close()
	entity.InitDataBase(db)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        entity.StarWarsSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={film(id:1){title}}'")
	http.ListenAndServe(":8080", nil)
}
