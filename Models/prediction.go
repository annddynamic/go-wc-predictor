package Models

import (
	"go-wc-predictor/Database"
	"log"
)

type Predictions []Prediction

type Prediction struct {
	Match               string `json:"match"`
	User                string `json:"user"`
	PredictionHomeScore int    `json:"predictionHomeScore"`
	PredictionAwayScore int    `json:"predictionAwayScore"`
}

func (prd Predictions) InsertPredictions(prediction Predictions) bool {
	for i := 0; i < len(prediction); i++ {
		prd.Insert(prediction[i])
	}
	return true
}

func (prd Predictions) Insert(prediction Prediction) {
	db := Database.GetInstance()

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
