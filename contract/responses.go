package contract

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetKeyValueRecordsResponse struct {
	Code     int       `json:"code"`
	Message  string    `json:"msg"`
	Records  []Record  `json:"records"`
}

type Record struct {
	Key        string               `json:"key"`
	CreatedAt  primitive.DateTime   `json:"createdAt"`
	TotalCount int                  `json:"totalCount"`
}

type SetKeyValueResponse KeyValuePair

type GetValueResponse KeyValuePair