package main

import (
	"go-wc-predictor/Models"
	"go-wc-predictor/server"
)

func main() {

	//db := Database.GetInstance()
	//_, err := db.Exec("INSERT INTO USERS VALUES(?,?);", "Andi", "Dika")
	//
	//if err != nil {
	//	fmt.Println("aa", err)
	//}
	prediction := Models.Prediction{"aa", "b", 1, 3}
	prediction.Insert(prediction)
	server.StartServer()

}
