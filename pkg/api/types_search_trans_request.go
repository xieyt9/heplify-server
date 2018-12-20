package api

type SearchTransRequestParamTZ struct {
	Value  int    `json:"value"`
	Name   string `json:"name"`
	Offset string `json:"offset"`
}

type SearchTransRequestParamLocNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchTransRequestParamLoc struct {
	Node []SearchTransRequestParamLocNode `json:"node"`
}

type SearchTransRequestParamSearch struct {
	ID     int      `json:"id"`
	CallID []string `json:"callid"`
	Uniq   bool     `json:"uniq"`
}

type SearchTransRequestParamTransaction struct {
	Call         bool `json:"call, omitempty"`
	Registration bool `json:"registration, omitempty"`
	Rest         bool `json:"rest, omitempty"`
}

type SearchTransRequestParam struct {
	Transaction SearchTransRequestParamTransaction `json:"transaction"`
	Search      SearchTransRequestParamSearch      `json:"search"`
	Location    SearchTransRequestParamLoc         `json:"location"`
	TimeZone    SearchTransRequestParamTZ          `json:"timezone"`
}

type SearchTransRequestTS struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type SearchTransRequest struct {
	Param     SearchTransRequestParam `json:"param"`
	TimeStamp SearchTransRequestTS    `json:"timestamp"`
}
