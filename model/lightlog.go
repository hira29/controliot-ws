package model

import "time"

//LogData = Create Log Data
type LogData struct {
	ID        string    `json:"_id" bson:"_id"`
	Lamp      int32     `json:"lamp" bson:"lamp"`
	Condition bool      `json:"condition" bson:"condition"`
	Time      time.Time `json:"time" bson:"time"`
}

//ReturnData = Create Log Data
type ReturnData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//RequestLog = RequestingLog
type RequestLog struct {
	Size    int64  `json:"size" example:"2"`
	Page    int64  `json:"page" example:"1"`
	Search  int32  `json:"search" example:"Admin"`
	OrderBy string `json:"order_by" example:"username"`
	Order   string `json:"order" example:"ASC"`
}

//Data = datastruct
type Data struct {
	Content  interface{} `json:"content" bson:"content"`
	PageInfo interface{} `json:"page_info" bson:"page_info"`
}

//Paging = pagingstruct
type Paging struct {
	Sort          bool   `json:"sort" bson:"sort"`
	SortBy        string `json:"sortBy" bson:"sortBy"`
	PageSize      int64  `json:"pageSize" bson:"pageSize"`
	PageNumber    int64  `json:"pageNumber" bson:"pageNumber"`
	TotalPages    int64  `json:"totalPages" bson:"totalPages"`
	TotalElements int64  `json:"totalElements" bson:"totalElements"`
}
