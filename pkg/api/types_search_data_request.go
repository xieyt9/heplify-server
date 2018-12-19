package api

type SearchDataRequestParamTZ struct {
	Value int `json:"value"`
	Name string `json:"name"`
	Offset string `json:"offset"`
}

type SearchDataRequestParamLocNode struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type SearchDataRequestParamLoc struct {
	Node []SearchDataRequestParamLocNode `json:"node"`
}

type SearchDataRequestParamSearch struct {
	ToUser string `json:"to_user"`
    Limit string `json:"limit"`
	FromUser string `json:"from_user"` 
	CallID string `json:"callid"`
}

type SearchDataRequestParamTransaction struct {
	Call bool `json:"call, omitempty"`
	Registration bool `json:"registration, omitempty"`
	Rest bool `json:"rest, omitempty"`
}

type SearchDataRequestParam struct{
	Transaction SearchDataRequestParamTransaction `json:"transaction"`
	Limit int `json:"limit"`
	Search SearchDataRequestParamSearch `json:"search"`
	Location SearchDataRequestParamLoc `json:"location"`
	TimeZone SearchDataRequestParamTZ `json:"timezone"`
}

type SearchDataRequestTS struct{
	From int64 `json:"from"`
	To int64 `json:"to"`
}

type SearchDataRequest struct{
	Param SearchDataRequestParam `json:"param"`
	TimeStamp SearchDataRequestTS `json:"timestamp"`
}