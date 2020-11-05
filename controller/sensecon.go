package controller

import (
	"controliot-ws/config"
	"controliot-ws/dao"
	"encoding/json"
	"github.com/gorilla/mux"
	//"log"
	"net/http"
)

//SensorSet = Data setting log database sensor
func SensorSet(w http.ResponseWriter, r *http.Request) {
	db := config.GetClient()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dao.SensorLog(params["dataSensor"], db))
}