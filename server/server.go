package server

import (
	"go-wc-predictor/client"
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

	if date == "" {
		_, err := w.Write([]byte("date required!"))
		if err != nil {
			log.Fatal("Error writing back: ", err)
		}
		return
	}

	response := srv.client.GetMatches(date)

	w.WriteHeader(200)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal("Error writing back: ", err)
	}

}

func StartServer() {

	server := &Server{client: client.NewClient()}

	http.HandleFunc("/api/matches", server.matches)
	log.Fatal(http.ListenAndServe("138.68.109.195:8080", nil))
}
