package contract

type GetCountsRequest struct {
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	MinCount   string   `json:"minCount"`
	MaxCount   string   `json:"maxCount"`
}

type KeyValuePair struct {
	Key    string    `json:"key"`
	Value  string    `json:"value"`
}

type SetKeyValueRequest KeyValuePair