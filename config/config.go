package config

import (
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

type auth struct {
	JwtSecret string `json:"jwtSecret"`
}

var (
	DBProperties   = make(map[string]database)
	AuthProperties *auth
)

func init() {
	log.Println("call config init()")

	setDBProperties()

	setAuthProperties()

}

func setAuthProperties() {
	f, err := os.Open("env/auth.json")

	checkErrForPanic(err)

	err = json.NewDecoder(f).Decode(&AuthProperties)

	checkErrForPanic(err)
}

func setDBProperties() {
	f, err := os.Open("env/database.json")

	checkErrForPanic(err)

	d := json.NewDecoder(f)

	for d.More() {
		var config database

		err := d.Decode(&config)

		checkErrForPanic(err)

		DBProperties[config.Kind] = config
	}
}

func checkErrForPanic(err error) {
	if err != nil {
		panic(err)
	}
}
