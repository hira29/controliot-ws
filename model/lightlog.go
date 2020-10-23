package model

import "time"

//Log = Create Log Data
type LogData struct {
	ID        string    `json:"_id" bson:"_id"`
	Lamp      int       `json:"lamp" bson:"lamp"`
	Condition bool      `json:"condition" bson:"condition"`
	Time      time.Time `json:"time" bson:"time"`
}

//ReturnData = Create Log Data
type ReturnData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
