package contract

type GetCountsResponse struct {
	Code     int       `json:"code"`
	Message  string    `json:"msg"`
	Records  []Record  `json:"records"`
}

type Record struct {
	Key        string   `json:"key"`
	CreatedAt  string   `json:"createdAt"`
	TotalCount int      `json:"totalCount"`
}

type SetKeyValueResponse KeyValuePair

type GetValueResponse KeyValuePair