package main

import (
	"BringMeFlyServer/database_mysql"
	"BringMeFlyServer/entity"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	/*
		db, err := bolt.Open("./database/swapi.db", 0600, nil)
		if err != nil {
			fmt.Println("error:", err)
		}
		defer db.Close()
	*/
	db := database_mysql.OpenDatabase("swapi")
	entity.InitDataBase(db)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        entity.StarWarsSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 5656")
	fmt.Println("Test with Get      : curl -g 'http://localhost:5656/graphql?query={film(id:1){title}}'")
	http.ListenAndServe(":5656", nil)
}
