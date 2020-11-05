package model

import "time"

//Log = Create Log Data
type LogSense struct {
	ID        string    `json:"_id" bson:"_id"`
	Sense     string    `json:"sense" bson:"sense"`
	Data	  int       `json:"data" bson:"data"`
	Time      time.Time `json:"time" bson:"time"`
}