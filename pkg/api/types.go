package api

import (
	"time"
)

//LoginSpec login spec
type LoginSpec struct {
	AuthName string `json:"authname,omitempty"`
	Auth     string `json:"auth,omitempty"`
	AuthID   string `json:"authID,omitempty"`
	Token    string `json:"token,omitempty"`
}

//Login login
type Login struct {
	Spec LoginSpec `json:"spec,omitempty"`
}

//Session request /api/v1/session
type Session struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SessionRespData resp data
type SessionRespData struct {
	UID       string    `json:"uid"`
	Username  string    `json:"username"`
	GID       string    `json:"gid"`
	GRP       string    `json:"grp"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	LastVisit time.Time `json:"lastvisit"`
}

//SessionResp response
type SessionResp struct {
	Status  int             `json:"status"`
	SID     string          `json:"sid"`
	Auth    string          `json:"auth"`
	Message string          `json:"message"`
	Data    SessionRespData `json:"data"`
}

type TimerRange struct {
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	Custom     time.Time `json:"custom"`
	CustomFrom time.Time `json:"customFrom"`
	CustomTo   time.Time `json:"customTo"`
}

type TimeZone struct {
	Value  string `json:"value"`
	Name   string `json:"name"`
	Offset string `json:"offset"`
}

type SearchConf struct {
	Limit  string `json:"limit"`
	Callid string `json:"callid"`
	ToUser string `json:"to_user"`
}

type Transaction struct {
}

type StoreResultType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type StoreResult struct {
	ResultType StoreResultType `json:"restype"`
}

type NodeItem struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type DBNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NodeConf struct {
	Node    NodeItem `json:"node"`
	DBNodes []DBNode `json:"dbnode"`
}

//ProfileStoreGetRespData profile store data
type ProfileStoreGetRespData struct {
	Range       TimerRange    `json:"timerrange"`
	TMZone      TimeZone      `json:"timezone"`
	Search      SearchConf    `json:"search"`
	Transaction []Transaction `json:"transaction"`
	Result      StoreResult   `json:"result"`
	Node        NodeConf      `json:"node"`
}

//ProfileStoreGetResp /api/v1/profile/store
type ProfileStoreGetResp struct {
	Status int                     `json:"status"`
	Data   ProfileStoreGetRespData `json:"data"`
}

type DashboardStoreDataItemSubItem struct {
	BoardID string `json:"boardid"`
	Name    string `json:"name"`
	Class   string `json:"class"`
}

type DashboardStoreDataItem struct {
	Name     string                          `json:"name"`
	Href     string                          `json:"href"`
	Class    string                          `json:"class"`
	RowClass string                          `json:"rowclass,omitempty"`
	ID       string                          `json:"id,omitempty"`
	SubItems []DashboardStoreDataItemSubItem `json:"subItems,omitempty"`
}

type DashboardNodeData struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

//DashboardStore /api/v1/dashboard/store
type DashboardStore struct {
	Status int                      `json:"status"`
	Data   []DashboardStoreDataItem `json:"data"`
}

//DashboardNode /api/v1/dashboard/node
type DashboardNode struct {
	Status int `json:"status"`
	Sid string 	`json:"sid"`
	Auth string `json:"auth"`
	Message string `json:"message"`
	Data []DashboardNodeData `json:"data"`
	Count int `json:"count"`
}

type ProfileStoreSearchData struct {
	
}

//ProfileStoreSearch /api/v1/profile/store/search
type ProfileStoreSearch struct {
	Status int `json:"status"`
	Data []ProfileStoreSearchData `json:"data"`
}