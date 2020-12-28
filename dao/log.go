package dao

import (
	"context"
	"controliot-ws/function"
	"controliot-ws/model"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetOn(led int32, db *mongo.Client) model.ReturnData {
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

func SetOff(led int32, db *mongo.Client) model.ReturnData {
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

func SensorLog(data string, num string, db *mongo.Client) model.ReturnData {
	var status bool
	var message string
	var returndata interface{}
	var sense model.LogSense

	dataSensor, error := strconv.Atoi(data)
	if error != nil {
		log.Println("Converting Data Error", error)
		status = false
		returndata = nil
		message = "error"
	} else {
		now := time.Now()
		now.Format(time.RFC3339)
		sense.ID = function.ToMd5(time.Now().String())
		sense.Sense = "US" + num
		sense.Data = dataSensor
		sense.Time = now

		col := db.Database("log").Collection("sense")
		insertResult, err := col.InsertOne(context.Background(), sense)
		if err != nil {
			log.Println("Error on inserting Log Document", err)
			status = false
			returndata = nil
			message = "error"
		}
		if insertResult.InsertedID == "" {
			status = false
			returndata = nil
			message = "error"
		} else {
			status = true
			returndata = sense
			message = "success"
		}
	}

	return model.ReturnData{Status: status, Data: returndata, Message: message}
}

func GetLightLog(inputData model.RequestLog, db *mongo.Client) model.ReturnData {
	var dataResult []model.LogData

	var sortIndex int
	findOptions := options.Find()
	if inputData.Size != 0 {
		findOptions.SetLimit(inputData.Size)
	}
	if inputData.Page != 0 {
		findOptions.SetSkip((inputData.Page - 1) * inputData.Size)
	}
	if inputData.OrderBy != "" && inputData.Order != "" {
		sortName := strings.ToLower(inputData.OrderBy)
		sortIndexName := strings.ToLower(inputData.Order)
		if sortIndexName == "asc" {
			sortIndex = 1
		} else {
			sortIndex = -1
		}

		findOptions.SetSort(bson.D{{sortName, sortIndex}})
	}

	searchOptions := bson.D{}
	if inputData.Search == 0 {
		searchOptions = bson.D{}
	} else {
		searchOptions = bson.D{{"lamp", inputData.Search}}
	}

	col := db.Database("log").Collection("light")

	cur, err := col.Find(context.TODO(), searchOptions, findOptions)
	if err != nil {
		log.Println("Error on finding all the documents", err)
		return model.ReturnData{Status: false, Data: nil, Message: "error on search: cant find index search"}
	}

	for cur.Next(context.TODO()) {
		var dataJson model.LogData

		err = cur.Decode(&dataJson)
		if err != nil {
			log.Println("Error on decoding the document", err)
			return model.ReturnData{Status: false, Data: nil, Message: "error on decoding"}
		}
		dataResult = append(dataResult, dataJson)
	}

	if dataResult == nil {
		cursor, error := col.Find(context.TODO(), bson.D{}, findOptions)
		if error != nil {
			log.Println("Error on finding all the documents", error)
			return model.ReturnData{Status: false, Data: nil, Message: "error on search: cant find index search"}
		}

		for cursor.Next(context.TODO()) {
			var data model.LogData

			err = cursor.Decode(&data)
			if err != nil {
				log.Println("2Error on decoding the document", err)
				return model.ReturnData{Status: false, Data: nil, Message: "error on decoding"}
			}

			dataResult = append(dataResult, data)
		}
		inputData.Search = 0
	}

	var Data model.Data
	var Paging model.Paging

	filter := bson.D{}

	if inputData.Search == 0 {
		filter = bson.D{}
	} else {
		filter = bson.D{{"lamp", inputData.Search}}
	}
	docsCount, err := col.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Println("Error on get Documents", err)
		return model.ReturnData{Status: false, Data: nil, Message: "error on get Documents"}
	}

	if inputData.Order == "ASC" {
		Paging.Sort = true
	} else if inputData.Order == "asc" {
		Paging.Sort = true
	} else {
		Paging.Sort = false
	}
	Paging.SortBy = inputData.OrderBy
	Paging.PageNumber = inputData.Page
	Paging.PageSize = inputData.Size
	data := math.Mod(float64(docsCount), float64(inputData.Size))
	dataset := docsCount / inputData.Size
	if dataset == 0 {
		Paging.TotalPages = dataset + 1
	} else {
		if data == 0 {
			Paging.TotalPages = dataset
		} else {
			Paging.TotalPages = dataset + 1
		}
	}

	Paging.TotalElements = docsCount

	Data.Content = dataResult
	Data.PageInfo = Paging

	return model.ReturnData{Status: true, Data: Data, Message: "success"}
}

func GetSenseLog(inputData model.RequestLog, db *mongo.Client) model.ReturnData {
	var dataResult []model.LogSense

	var sortIndex int
	findOptions := options.Find()
	if inputData.Size != 0 {
		findOptions.SetLimit(inputData.Size)
	}
	if inputData.Page != 0 {
		findOptions.SetSkip((inputData.Page - 1) * inputData.Size)
	}
	if inputData.OrderBy != "" && inputData.Order != "" {
		sortName := strings.ToLower(inputData.OrderBy)
		sortIndexName := strings.ToLower(inputData.Order)
		if sortIndexName == "asc" {
			sortIndex = 1
		} else {
			sortIndex = -1
		}

		findOptions.SetSort(bson.D{{sortName, sortIndex}})
	}

	col := db.Database("log").Collection("sense")

	cur, err := col.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Println("Error on finding all the documents", err)
		return model.ReturnData{Status: false, Data: nil, Message: "error on search: cant find index search"}
	}

	for cur.Next(context.TODO()) {
		var dataJson model.LogSense

		err = cur.Decode(&dataJson)
		if err != nil {
			log.Println("Error on decoding the document", err)
			return model.ReturnData{Status: false, Data: nil, Message: "error on decoding"}
		}
		dataResult = append(dataResult, dataJson)
	}

	var Data model.Data
	var Paging model.Paging

	filter := bson.D{}

	docsCount, err := col.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Println("Error on get Documents", err)
		return model.ReturnData{Status: false, Data: nil, Message: "error on get Documents"}
	}

	if inputData.Order == "ASC" {
		Paging.Sort = true
	} else if inputData.Order == "asc" {
		Paging.Sort = true
	} else {
		Paging.Sort = false
	}
	Paging.SortBy = inputData.OrderBy
	Paging.PageNumber = inputData.Page
	Paging.PageSize = inputData.Size
	data := math.Mod(float64(docsCount), float64(inputData.Size))
	dataset := docsCount / inputData.Size
	if dataset == 0 {
		Paging.TotalPages = dataset + 1
	} else {
		if data == 0 {
			Paging.TotalPages = dataset
		} else {
			Paging.TotalPages = dataset + 1
		}
	}

	Paging.TotalElements = docsCount

	Data.Content = dataResult
	Data.PageInfo = Paging

	return model.ReturnData{Status: true, Data: Data, Message: "success"}
}
