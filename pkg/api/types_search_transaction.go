package api

type TransactionDataCalldata struct {
	ID               string `json:"id" gorm:"id"`
	Protocol         string `json:"protocol"`
	Method           string `json:"method" gorm:"method"`
	ReplyReason      string `json:"reply_reason" gorm:"reply_reason"`
	SourcePort       string `json:"src_port" gorm:"source_port"`
	DestinationPort  string `json:"dst_port" gorm:"destination_port"`
	Trans            string `json:"trans"`
	CallID           string `json:"callid" gorm:"callid"`
	Node             string `json:"node" gorm:"node"`
	DbNode           string `json:"dbnode" gorm:"dbnode"`
	MicroTs          string `json:"micro_ts" gorm:"micro_ts"`
	RuriUser         string `json:"ruri_user" gorm:"ruri_user"`
	SourceAlias      string `json:"source_alias" gorm:"source_alias"`
	DestinationAlias string `json:"destination_alias" gorm:"destination_alias"`
	SourceIp         string `json:"source_ip" gorm:"source_ip"`
	DestinationIp    string `json:"destination_ip" gorm:"destination_ip"`
	SrcID            string `json:"src_id"`
	DstID            string `json:"dst_id"`
	MilliTs          int    `json:"milli_ts"`
	MethodText       string `json:"method_text"`
	MsgColor         string `json:"msg_color"`
	Destination      int    `json:"destination"`
}

type TransactionDataHostsContent struct {
	Position int              `json:"position"`
	IsStp    int              `json:"is_stp"`
	Hosts    []map[string]int `json:"hosts"`
}

type TransactionDataInfoCallId struct {
}

type TransactionDataInfo struct {
	CallID     []TransactionDataInfoCallId `json:"callid"`
	TotDur     string                      `json:"totdur"`
	StatusCall int                         `json:"statuscall"`
}

type TransactionDataUAC struct {
	Image string `json:"image"`
	Agent string `json:"agent"`
}

type TransactionDataRTPInfo struct {
}

type TransactionData struct {
	Info     TransactionDataInfo                    `json:"info"`
	HostsMap map[string]TransactionDataHostsContent `json:"hosts"`
	UACMap   map[string]TransactionDataUAC          `json:"uac"`
	RTPInfo  []TransactionDataRTPInfo               `json:"rtpino"`
	CallData []TransactionDataCalldata              `json:"calldata"`
	Count    int                                    `json:"count"`
}

//SearchTransaction /api/v1/search/transaction
type SearchTransaction struct {
	Status  int             `json:"status"`
	Sid     string          `json:"sid"`
	Auth    string          `json:"auth"`
	Message string          `json:"message"`
	Data    TransactionData `json:"data"`
	Count   int             `json:"count"`
}
