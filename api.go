
package main

import (
	"encoding/json"
	m "github.com/gorilla/mux" // can type alias for the import
	"log"
	"net/http"
	time2 "time"
)

type Time struct {
	Now string `json:"now,omitempty"`
	Message string `json:"m"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {

	time := Time{
		Now: time2.Now().String(),
		Message: "hello",
	}

	log.Printf("Get request from %v", r.RemoteAddr)

	//w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(time)
}

func CreateTime(w http.ResponseWriter, r *http.Request) {

	params :=m.Vars(r)

	log.Printf("%v", params)
	var time Time

	json.NewDecoder(r.Body).Decode(&time)

	log.Print("%v", time)

}


func main() {
	router := m.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/", CreateTime).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
