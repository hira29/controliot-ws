package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tmpControl = false

func setOn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := make(map[string]interface{})
	Data["message"] = "Lampu On"
	Data["status"] = true
	Data["data"] = togleLamp(true)
	json.NewEncoder(w).Encode(Data)
}

func setOff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := make(map[string]interface{})
	Data["message"] = "Lampu Off"
	Data["status"] = true
	Data["data"] = togleLamp(false)
	json.NewEncoder(w).Encode(Data)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := make(map[string]interface{})
	if tmpControl == true {
		Data["message"] = "Lampu On"
		Data["status"] = true
		Data["data"] = tmpControl
	} else {
		Data["message"] = "Lampu Off"
		Data["status"] = true
		Data["data"] = tmpControl
	}
	json.NewEncoder(w).Encode(Data)
}

func togleLamp(control bool) bool {
	tmpControl = control
	return tmpControl
}

func main() {
	//port := os.Getenv("PORT")
	port := "8080"
	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{
		"X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token",
	})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	r.HandleFunc("/on", setOn).Methods("GET")
	r.HandleFunc("/off", setOff).Methods("GET")
	r.HandleFunc("/status", getStatus).Methods("GET")

	log.Println("API STARTED!")
	log.Println(port)
	_ = http.ListenAndServe(":"+port, handlers.CORS(headers, origins, methods)(r))
}
