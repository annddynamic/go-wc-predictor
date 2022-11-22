package Models

import (
	"go-wc-predictor/Database"
	"log"
)

type Prediction struct {
	Match               string `json:"match"`
	User                string `json:"user"`
	PredictionHomeScore int    `json:"predictionHomeScore"`
	PredictionAwayScore int    `json:"predictionAwayScore"`
}

func (prd Prediction) Insert(prediction Prediction) bool {
	db := Database.GetInstance()

	stmt, err := db.Prepare("INSERT INTO PREDICTIONS VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(prediction.Match, prediction.User, prediction.PredictionHomeScore, prediction.PredictionAwayScore)
	if err != nil {
		log.Fatal("Error inserting into predictions", err)
		return false
	}

	return true
}
