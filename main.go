package main

import (
	"fmt"
	"go-wc-predictor/Database"
	"go-wc-predictor/server"
)

func main() {

	db := Database.GetInstance()
	_, err := db.Exec("INSERT INTO USERS VALUES(?,?);", "Andi", "Dika")

	if err != nil {
		fmt.Println("aa", err)
	}
	server.StartServer()

}
