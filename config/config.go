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

type auth struct {
	JwtSecret string `json:"jwtSecret"`
}

var DBProperties = make(map[string]database)
var AuthProperties auth

type key int

const (
	MySQL key = iota
	Logger
	JWTInfo
)

func init() {
	log.Println("call config init()")

	setDBProperties()

	setAuthProperties()

}

func setAuthProperties() {
	f, err := os.Open("config/auth.json")

	checkErrForPanic(err)

	err = json.NewDecoder(f).Decode(&AuthProperties)

	checkErrForPanic(err)
}

func setDBProperties() {
	f, err := os.Open("config/database.json")

	checkErrForPanic(err)

	d := json.NewDecoder(f)

	for d.More() {
		var config database

		err := d.Decode(&config)

		checkErrForPanic(err)

		DBProperties[config.Kind] = config
	}
}

func NewMySQL() *sql.DB {
	con := DBProperties["mysql"]

	url := con.Username + ":" + con.Password + "@tcp(" + con.Address + ":" + con.Port + ")/" + con.Name + "?parseTime=true"

	db, err := sql.Open(con.Kind, url)

	checkErrForPanic(err)

	err = db.Ping()

	checkErrForPanic(err)

	// 항상 연결되어있는 커넥션 풀 갯수 설정
	db.SetMaxIdleConns(5)
	// 최대로 연결할 수 있는 커넥션 갯수 설정
	db.SetMaxOpenConns(10)

	return db
}

func checkErrForPanic(err error) {
	if err != nil {
		panic(err)
	}
}
