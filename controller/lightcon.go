package controller

import (
	"controliot-ws/config"
	"controliot-ws/dao"
	"encoding/json"
	"log"
	"net/http"
)

//TmpControl1  = Controll Temp Data Save
var TmpControl1 = false

//TmpControl2 = Controll Temp Data Save
var TmpControl2 = false

// SetOn1 = Setting On
func SetOn1(w http.ResponseWriter, r *http.Request) {
	db := config.GetClient()
	log.Println("Connected")
	w.Header().Set("Content-Type", "application/json")
	//var ret int
	_ = TogleLamp(true, 1)
	json.NewEncoder(w).Encode(dao.SetOn(1, db))
}

// SetOff1 = Controlling Off
func SetOff1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.GetClient()
	//var ret int
	_ = TogleLamp(false, 1)
	json.NewEncoder(w).Encode(dao.SetOff(1, db))
}

// SetOn2 = Setting On
func SetOn2(w http.ResponseWriter, r *http.Request) {
	db := config.GetClient()
	w.Header().Set("Content-Type", "application/json")
	//var ret int
	_ = TogleLamp(true, 2)
	json.NewEncoder(w).Encode(dao.SetOn(2, db))
}

// SetOff2 = Controlling Off
func SetOff2(w http.ResponseWriter, r *http.Request) {
	db := config.GetClient()
	w.Header().Set("Content-Type", "application/json")
	//var ret int
	_ = TogleLamp(false, 2)
	json.NewEncoder(w).Encode(dao.SetOff(2, db))
}

//GetStatus1 = Getting Status
func GetStatus1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := make(map[string]interface{})
	if TmpControl1 == true {
		Data["message"] = "Lampu On"
		Data["status"] = true
		Data["data"] = TmpControl1
	} else {
		Data["message"] = "Lampu Off"
		Data["status"] = true
		Data["data"] = TmpControl1
	}
	json.NewEncoder(w).Encode(Data)
}

//GetStatus2 = Getting Status
func GetStatus2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := make(map[string]interface{})
	if TmpControl2 == true {
		Data["message"] = "Lampu On"
		Data["status"] = true
		Data["data"] = TmpControl2
	} else {
		Data["message"] = "Lampu Off"
		Data["status"] = true
		Data["data"] = TmpControl2
	}
	json.NewEncoder(w).Encode(Data)
}

//TogleLamp = Togling Lamp
func TogleLamp(control bool, i int) int {
	if i == 1 {
		TmpControl1 = control
	} else if i == 2 {
		TmpControl2 = control
	}
	return i
}
