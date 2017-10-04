package main

import (
	"encoding/json"
	"log"
        "net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting heu-api")
        readConfigFile(configFile)
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET", "POST")

	http.Handle("/", r)
	http.ListenAndServe(":80", r)

}

func index(w http.ResponseWriter, r *http.Request) {
	s := IndexResponse{}

	switch r.Method {
	case "GET":
		s = IndexResponse{Success: 1, Message: "Hello ! Index test with GET method"}
	case "POST":
		s = IndexResponse{Success: 1, Message: "Hello ! Index test with POST method"}
	default:
		s = IndexResponse{Success: 1, Message: "Hello ! Index test with not allowed method"}
	}

	response, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(response)
}
