package config

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
}

var dbProperties = make(map[string]database)

func init() {
	log.Println("call config init()")

	f, err := os.Open("config/database.json")

	checkErrForPanic(err)

	d := json.NewDecoder(f)

	for d.More() {
		var config database

		err := d.Decode(&config)

		checkErrForPanic(err)

		dbProperties[config.Kind] = config
	}

}

func NewMySQL() *sql.DB {
	con := dbProperties["mysql"]

	url := con.Username + ":" + con.Password + "@tcp(" + con.Address + ":" + con.Port + ")/" + con.Name

	db, err := sql.Open(con.Kind, url)

	checkErrForPanic(err)

	err = db.Ping()

	checkErrForPanic(err)

	return db
}

func checkErrForPanic(err error) {
	if err != nil {
		panic(err)
	}
}
