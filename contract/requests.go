package contract

type GetKeyValueRecordsRequest struct {
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	MinCount   int      `json:"minCount"`
	MaxCount   int      `json:"maxCount"`
}

type KeyValuePair struct {
	Key    string    `json:"key"`
	Value  string    `json:"value"`
}

type SetKeyValueRequest KeyValuePair