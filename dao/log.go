package dao

import (
	"context"
	"controliot-ws/function"
	"controliot-ws/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetOn(led int, db *mongo.Client) model.ReturnData {
	var status bool
	var message string
	var data interface{}
	var dataReturn model.LogData

	now := time.Now()
	now.Format(time.RFC3339)
	dataReturn.ID = function.ToMd5(time.Now().String())
	dataReturn.Lamp = led
	dataReturn.Condition = true
	dataReturn.Time = now

	col := db.Database("log").Collection("light")
	insertResult, err := col.InsertOne(context.Background(), dataReturn)
	if err != nil {
		log.Println("Error on inserting Log Document", err)
		status = false
		data = nil
		message = "error"
	}
	if insertResult.InsertedID == "" {
		status = false
		data = nil
		message = "error"
	} else {
		status = true
		data = dataReturn
		message = "success"
	}
	return model.ReturnData{Status: status, Data: data, Message: message}
}

func SetOff(led int, db *mongo.Client) model.ReturnData {
	var status bool
	var message string
	var data interface{}
	var dataReturn model.LogData

	now := time.Now()
	now.Format(time.RFC3339)
	dataReturn.ID = function.ToMd5(time.Now().String())
	dataReturn.Lamp = led
	dataReturn.Condition = false
	dataReturn.Time = now

	col := db.Database("log").Collection("light")
	insertResult, err := col.InsertOne(context.Background(), dataReturn)
	if err != nil {
		log.Println("Error on inserting Log Document", err)
		status = false
		data = nil
		message = "error"
	}
	if insertResult.InsertedID == "" {
		status = false
		data = nil
		message = "error"
	} else {
		status = true
		data = dataReturn
		message = "success"
	}
	return model.ReturnData{Status: status, Data: data, Message: message}
}
