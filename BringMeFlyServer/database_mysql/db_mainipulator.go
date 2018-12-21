package database_mysql

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTable(db *sql.DB, table string) {
	sql := `
		CREATE TABLE IF NOT EXISTS` + table + ` (
			id INT,
			data JSON,
			PRIMARY KEY(id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	smt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	smt.Exec()
}
func OpenDatabase(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:RootPassword1!@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	return db
}
func Search_by_kind_and_id(db *sql.DB, table string, selectId int) string {
	tx, _ := db.Begin()
	defer tx.Commit()
	rows, _ := tx.Query("SELECT * FROM " + table)
	defer rows.Close()
	for rows.Next() {
		var data []byte
		var id int
		if err := rows.Scan(&id, &data); err != nil {
			log.Fatal(err)
		}
		if id == selectId {
			return string(data)
		}
	}
	return ""
}

func Unmarshal_fields(s string) string {
	var _json map[string]interface{}
	json.Unmarshal([]byte(s), &_json)
	//fmt.Println(_json)
	temp, _ := json.Marshal(_json["fields"])
	//fmt.Println(temp)
	return string(temp)
	//return value
}
