package main

import (
	"anyweb/config"
)

func main() {

	db := config.NewMySQL()

	defer db.Close()

}
