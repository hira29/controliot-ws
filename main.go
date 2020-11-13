package main

import (
	"controliot-ws/controller"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//port := os.Getenv("PORT")
	port := "8080"
	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{
		"X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token",
	})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	r.HandleFunc("/on/1", controller.SetOn1).Methods("GET")
	r.HandleFunc("/off/1", controller.SetOff1).Methods("GET")
	r.HandleFunc("/on/2", controller.SetOn2).Methods("GET")
	r.HandleFunc("/off/2", controller.SetOff2).Methods("GET")
	r.HandleFunc("/status/1", controller.GetStatus1).Methods("GET")
	r.HandleFunc("/status/2", controller.GetStatus2).Methods("GET")
	r.HandleFunc("/sensor/{dataSensor}", controller.SensorSet).Methods("GET")
	r.HandleFunc("/log/light", controller.GetLightLog).Methods("POST")
	r.HandleFunc("/log/sense", controller.GetSenseLog).Methods("POST")

	log.Println("API STARTED!")
	log.Println(port)
	_ = http.ListenAndServe(":"+port, handlers.CORS(headers, origins, methods)(r))
}
