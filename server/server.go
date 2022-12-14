package server

import (
	"encoding/json"
	"fmt"
	"go-wc-predictor/Models"
	"go-wc-predictor/client"
	"io"
	"log"
	"net/http"
)

type Server struct {
	client client.Client
}

func (srv *Server) allowCors(w http.ResponseWriter) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func (srv *Server) matches(w http.ResponseWriter, r *http.Request) {

	srv.allowCors(w)

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
		return
	}
	return
}

// post
func (srv *Server) predict(w http.ResponseWriter, r *http.Request) {

	srv.allowCors(w)

	switch r.Method {
	case http.MethodPost:
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("server: could not read request body: %s\n", err)
		}

		prediction := Models.Predictions{}
		err = json.Unmarshal(reqBody, &prediction)

		if err != nil {
			fmt.Println("Error unmarshaling json: ", err)
			return
		}

		if prediction.InsertPredictions(prediction) {
			w.WriteHeader(200)
			srv.respond(w, "Success")
			return
		}
		w.WriteHeader(500)
		srv.respond(w, "Internal server error")
		return
	case "OPTIONS":
		w.WriteHeader(200)
	default:
		http.Error(w, fmt.Sprintf("method %s is not allowed", r.Method), http.StatusMethodNotAllowed)
	}

}

func (srv *Server) respond(w http.ResponseWriter, message string) {
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)

	if err != nil {
		log.Fatalf("Error writing back: %s", err)
	}
}

func StartServer() {

	server := &Server{client: client.NewClient()}

	http.HandleFunc("/api/matches", server.matches)
	http.HandleFunc("/api/predict", server.predict)
	fmt.Println("Starting server at: localhost:5000")
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
