package server

import (
	"fmt"
	"go-wc-predictor/client"
	"io"
	"log"
	"net/http"
)

type Server struct {
	client client.Client
}

func (srv *Server) matches(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	query := r.URL.Query()
	date := query.Get("date") //filters="color"
	fmt.Println()
	if date == "" {
		_, err := w.Write([]byte("date required!"))
		if err != nil {
			log.Fatal("Error writing back: ", err)
		}
		return
	}

	response := srv.client.GetMatches(date)
	//
	w.WriteHeader(200)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal("Error writing back: ", err)
	}

}

// post
func (srv *Server) predict(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}

	/*
		reqBody should be:
		[
			{"match":"ArgentinaSaudi Arabia","user":"vigan","predictedHomeScore":1,"predictedAwayScore":2},
			{"match":"MexicoPoland","user":"vigan","predictedHomeScore":3,"predictedAwayScore":4},
			{"match":"DenmarkTunisia","user":"vigan","predictedHomeScore":0,"predictedAwayScore":0},
			{"match":"FranceAustralia","user":"vigan","predictedHomeScore":0,"predictedAwayScore":0}
		]

		tabela n sql lite should be:


	*/
	//
	fmt.Println(reqBody)
	//w.WriteHeader(200)
	//_, err := w.Write(response)
	//if err != nil {
	//	log.Fatal("Error writing back: ", err)
	//}

}

func StartServer() {

	server := &Server{client: client.NewClient()}

	http.HandleFunc("/api/matches", server.matches)
	http.HandleFunc("/api/predict", server.predict)
	//http.HandleFunc("/api/predict", server.predict)
	fmt.Println("Starting server at: 138.68.109.195:443")
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
