package controller

import (
	"context"
	"controliot-ws/config"
	"controliot-ws/dao"
	"controliot-ws/model"
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
	json.NewEncoder(w).Encode(dao.SensorLog(params["dataSensor"], params["numSensor"], db))
	defer db.Disconnect(context.Background())
}

//GetSenseLog = Getting Log
func GetSenseLog(w http.ResponseWriter, r *http.Request) {
	db := config.GetClient()
	var inputRequest model.RequestLog
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&inputRequest)
	json.NewEncoder(w).Encode(dao.GetSenseLog(inputRequest, db))
	defer db.Disconnect(context.Background())
}
