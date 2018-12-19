package api

import (
    utils "github.com/sipcapture/heplify-server/pkg/utils"
)

type SIPData struct {
    Id               string         `json:"id" gorm:"id"`
    Date             utils.JsonTime `json:"date" gorm:"date"`
    MilliTs          string         `json:"milli_ts"`
    MicroTs          string         `json:"micro_ts" gorm:"micro_ts"`
    Method           string         `json:"method" gorm:"method"`
    ReplyReason      string         `json:"reply_reason" gorm:"reply_reason"`
    Ruri             string         `json:"ruri" gorm:"ruri"`
    RuriUser         string         `json:"ruri_user" gorm:"ruri_user"`
    RuriDomain       string         `json:"ruri_domain" gorm:"ruri_domain"`
    FromUser         string         `json:"from_user" gorm:"from_user"`
    FromDomain       string         `json:"from_domain" gorm:"from_domain"`
    FromTag          string         `json:"from_tag" gorm:"from_tag"`
    ToUser           string         `json:"to_user" gorm:"to_user"`
    ToDomain         string         `json:"to_domain" gorm:"to_domain"`
    ToTag            string         `json:"to_tag" gorm:"to_tag"`
    PidUser          string         `json:"pid_user" gorm:"pid_user"`
    ContactUser      string         `json:"contact_user" gorm:"contact_user"`
    AuthUser         string         `json:"auth_user" gorm:"auth_user"`
    Callid           string         `json:"callid" gorm:"callid"`
    CallidAleg       string         `json:"callid_aleg" gorm:"callid_aleg"`
    Via_1            string         `json:"via_1" gorm:"via_1"`
    Via_1_Branch     string         `json:"via_1_branch" gorm:"via_1_branch"`
    Cseq             string         `json:"cseq" gorm:"cseq"`
    Diversion        string         `json:"diversion" gorm:"diversion"`
    Reason           string         `json:"reason" gorm:"reason"`
    ContentType      string         `json:"content_type" gorm:"content_type"`
    Auth             string         `json:"auth" gorm:"auth"`
    UserAgent        string         `json:"user_agent" gorm:"user_agent"`
    SourceIp         string         `json:"source_ip" gorm:"source_ip"`
    SourcePort       string         `json:"source_port" gorm:"source_port"`
    DestinationIp    string         `json:"destination_ip" gorm:"destination_ip"`
    DestinationPort  string         `json:"destination_port" gorm:"destination_port"`
    ContactIp        string         `json:"contact_ip" gorm:"contact_ip"`
    ContactPort      string         `json:"contact_port" gorm:"contact_port"`
    OriginatorIp     string         `json:"originator_ip" gorm:"originator_ip"`
    OriginatorPort   string         `json:"originator_port" gorm:"originator_port"`
    Expires          string         `json:"expires, omitempty" gorm:"expires"`
    CorrelationId    string         `json:"correlation_id" gorm:"correlation_id"`
    Proto            string         `json:"proto" gorm:"proto"`
    Family           string         `json:"family" gorm:"family"`
    RtpStat          string         `json:"rtp_stat" gorm:"rtp_stat"`
    Type             string         `json:"type" gorm:"type"`
    Node             string         `json:"node" gorm:"node"`
    CustomField1     string         `json:"custom_field1" gorm:"custom_field1"`
    CustomField2     string         `json:"custom_field2" gorm:"custom_field2"`
    CustomField3     string         `json:"custom_field3" gorm:"custom_field3"`
    Trans            string         `json:"trans" gorm:"trans"`
    Dbnode           string         `json:"dbnode"`
    SourceAlias      string         `json:"source_alias" gorm:"source_alias"`
    DestinationAlias string         `json:"destination_alias" gorm:"destination_alias"`
    Msg              string         `json:"msg" gorm:"msg"`
}

//SearchData /api/v1/search/data
type SearchData struct {
    Status  int       `json:"status"`
    Sid     string    `json:"sid"`
    Auth    string    `json:"auth"`
    Message string    `json:"message"`
    Data    []SIPData `json:"data"`
}
