package api

type Text struct {
	TextHead string
	TextMsg  string
}

type TextData struct {
	MicroTs         int64  `gorm:"micro_ts"`
	SourceIp        string `gorm:"source_ip"`
	SourcePort      string `gorm:"source_port"`
	DestinationIp   string `gorm:"destination_ip"`
	DestinationPort string `gorm:"destination_port"`
	Msg             string `gorm:"msg"`
}
