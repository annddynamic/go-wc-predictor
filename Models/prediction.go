package Models

import (
	"go-wc-predictor/Database"
	"log"
)

type Prediction struct {
	Match               string
	User                string
	PredictionHomeScore int
	PredictionAwayScore int
}

func (prd Prediction) Insert(prediction Prediction) {
	db := Database.GetInstance()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO PREDICTIONS VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(prediction.Match, prediction.User, prediction.PredictionHomeScore, prediction.PredictionAwayScore)
	if err != nil {
		log.Fatal("Error inserting into predictions", err)
	}
}
