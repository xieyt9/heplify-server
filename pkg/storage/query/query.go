package query

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/storage/helper"
	"github.com/sipcapture/heplify-server/pkg/storage/mysqls"
	utils "github.com/sipcapture/heplify-server/pkg/utils"
)

const (
	logsCaputureTable   = "logs_capture"
	reportCaputureTable = "report_capture"

	//maintain these table
	isupCaputureAllTable    = "isup_capture_all"
	rtcpCaputureAllTable    = "rtcp_capture_all"
	sipCaputureCallAllTable = "sip_capture_call"
	sipCaputureRegAllTable  = "sip_capture_registration"
	sipCaputureRestAllTable = "sip_capture_rest"
)

func QuerySIPCaptureCall(searchdatareq api.SearchDataRequest) ([]api.SIPData, error) {

	store := helper.GetMysqlSotrage()
	handle, ok := store.(*mysqls.Store)
	if !ok {
		return nil, fmt.Errorf("need mysql handle")
	}

	dbhandle := handle.Client

	var sipdata []api.SIPData
	timefrom := searchdatareq.TimeStamp.From
	timeto := searchdatareq.TimeStamp.To
	fromparam := searchdatareq.Param.Search.FromUser
	//toparam := searchdatareq.Param.Search.ToUser
	callid := searchdatareq.Param.Search.CallID

	fromsplitcode, fromsplitresult := utils.SplitParam(fromparam)
	//tosplitcode, tosplitresult := utils.SplitParam(toparam)

	//	fromdate := utils.TsToDT(timefrom)
	//	todate := utils.TsToDT(timeto)
	fromdate := utils.TsToDTUTC(timefrom)
	todate := utils.TsToDTUTC(timeto)
	fromdateutc := utils.TsToDTUTC(timefrom)
	todateutc := utils.TsToDTUTC(timeto)

	// 得到所需查询的数据库表
	var tablename []string
	tablenametmp := ""
	datetmp := fromdateutc
	for i := 0; datetmp.Before(todateutc); {
		if searchdatareq.Param.Transaction.Call {
			tablenametmp = fmt.Sprintf("%s_%s", sipCaputureCallAllTable, utils.DTToDBNSuf(datetmp))
			tablename = append(tablename, tablenametmp)
			i++
		}
		if searchdatareq.Param.Transaction.Registration {
			tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRegAllTable, utils.DTToDBNSuf(datetmp))
			tablename = append(tablename, tablenametmp)
			i++
		}
		if searchdatareq.Param.Transaction.Rest {
			tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRestAllTable, utils.DTToDBNSuf(datetmp))
			tablename = append(tablename, tablenametmp)
			i++
		}
		datetmp = datetmp.AddDate(0, 0, 1)
	}

	// 对所需数据库表进行查询
	for idx := range tablename {
		var sipdatatmp []api.SIPData
		dbhandletmp := dbhandle.Table(tablename[idx])
		dbhandletmp = dbhandletmp.Where("date BETWEEN ? AND ?", fromdate, todate)
		if fromsplitcode != 0 {
			if fromsplitcode == 1 {
				dbhandletmp = dbhandletmp.Where("from_user = ?", fromsplitresult[0])
			} else if fromsplitcode == 2 {
				dbhandletmp = dbhandletmp.Where("source_ip = ?", fromsplitresult[0])
			} else if fromsplitcode == 3 {
				dbhandletmp = dbhandletmp.Where("from_user = ?", fromsplitresult[0])
				dbhandletmp = dbhandletmp.Where("source_ip = ?", fromsplitresult[1])
			}
		}
		/*if tosplitcode != 0 && strings.Contains(tablename[idx], sipCaputureCallAllTable) {
			if tosplitcode == 1 {
				dbhandletmp = dbhandletmp.Where("to_user = ?", tosplitresult[0])
			} else if tosplitcode == 2 {
				dbhandletmp = dbhandletmp.Where("destination_ip = ?", tosplitresult[0])
			} else if tosplitcode == 3 {
				dbhandletmp = dbhandletmp.Where("to_user = ?", tosplitresult[0])
				dbhandletmp = dbhandletmp.Where("destination_ip = ?", tosplitresult[1])
			}
		}*/
		if callid != "" && strings.Contains(tablename[idx], sipCaputureCallAllTable) {
			dbhandletmp = dbhandletmp.Where("callid = ?", callid)
		}
		dbhandletmp = dbhandletmp.Limit(250)
		//dbhandletmp = dbhandletmp.Order("id")
		dbhandletmp.Find(&sipdatatmp)
		if len(sipdatatmp) != 0 {
			sipdata = append(sipdata, sipdatatmp...)
		}
	}
	return sipdata, nil
}

func QueryMethod(searchmethodreq api.SearchMethodReq) ([]api.SearchMethodData, error) {
	store := helper.GetMysqlSotrage()
	handle, ok := store.(*mysqls.Store)
	if !ok {
		return nil, fmt.Errorf("need mysql handle")
	}

	dbhandle := handle.Client

	var methoddata []api.SearchMethodData
	timefrom := searchmethodreq.TimeStamp.From
	timeto := searchmethodreq.TimeStamp.To
	id := searchmethodreq.Param.Search.ID
	callid := searchmethodreq.Param.Search.CallID
	fromdateutc := utils.TsToDTUTC(timefrom)

	var tablename []string
	tablenametmp := ""
	datetmp := fromdateutc

	tablenametmp = fmt.Sprintf("%s_%s", sipCaputureCallAllTable, utils.DTToDBNSuf(datetmp))
	tablename = append(tablename, tablenametmp)

	tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRegAllTable, utils.DTToDBNSuf(datetmp))
	tablename = append(tablename, tablenametmp)

	tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRestAllTable, utils.DTToDBNSuf(datetmp))
	tablename = append(tablename, tablenametmp)

	for idx := range tablename {
		var methoddatatmp []api.SearchMethodData
		dbhandletmp := dbhandle.Table(tablename[idx])
		dbhandletmp = dbhandletmp.Where("micro_ts BETWEEN ? AND ?", timefrom*1000, timeto*1000+999)
		dbhandletmp = dbhandletmp.Where("id = ?", id)
		dbhandletmp = dbhandletmp.Where("callid = ?", callid)
		dbhandletmp.Find(&methoddatatmp)
		if len(methoddatatmp) != 0 {
			methoddata = append(methoddata, methoddatatmp...)
		}
	}
	return methoddata, nil
}

func QueryTransactionData(searchtransreq api.SearchTransRequest) (api.TransactionData, error) {
	store := helper.GetMysqlSotrage()
	handle, ok := store.(*mysqls.Store)
	if !ok {
		fmt.Println("get handle not ok")
	}
	dbhandle := handle.Client

	var transdata api.TransactionData
	var calldata []api.TransactionDataCalldata

	timefrom := searchtransreq.TimeStamp.From
	timeto := searchtransreq.TimeStamp.To
	callid := searchtransreq.Param.Search.CallID

	fromdateutc := utils.TsToDTUTC(timefrom)
	todateutc := utils.TsToDTUTC(timeto)

	// 得到所需查询的数据库表
	var tablename []string
	tablenametmp := ""
	datetmp := fromdateutc
	for i := 0; datetmp.Before(todateutc); {
		tablenametmp = fmt.Sprintf("%s_%s", sipCaputureCallAllTable, utils.DTToDBNSuf(datetmp))
		tablename = append(tablename, tablenametmp)
		if tablenametmp != "" {
			i++
		}

		tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRegAllTable, utils.DTToDBNSuf(datetmp))
		tablename = append(tablename, tablenametmp)
		if tablenametmp != "" {
			i++
		}

		tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRestAllTable, utils.DTToDBNSuf(datetmp))
		tablename = append(tablename, tablenametmp)
		if tablenametmp != "" {
			i++
		}
		datetmp = datetmp.AddDate(0, 0, 1)
	}

	// 对所需数据库表进行查询
	for idx := range tablename {
		var calldatatmp []api.TransactionDataCalldata
		dbhandletmp := dbhandle.Table(tablename[idx])
		dbhandletmp = dbhandletmp.Where("micro_ts BETWEEN ? AND ?", timefrom*1000, timeto*1000+999)
		for idx := range callid {
			dbhandletmp = dbhandletmp.Where("callid = ?", callid[idx])
		}
		dbhandletmp = dbhandletmp.Order("micro_ts")
		dbhandletmp = dbhandletmp.Find(&calldatatmp)
		if !dbhandletmp.RecordNotFound() {
			if strings.Contains(tablename[idx], "call") {
				for idx2 := range calldatatmp {
					calldatatmp[idx2].Trans = "call"
				}
			} else if strings.Contains(tablename[idx], "registration") {
				for idx2 := range calldatatmp {
					calldatatmp[idx2].Trans = "registration"
				}
			} else {
				for idx2 := range calldatatmp {
					calldatatmp[idx2].Trans = "rest"
				}
			}
			calldata = append(calldata, calldatatmp...)
		}
	}

	//生成UAC
	type SIPUAC struct {
		SourceIp   string `json:"source_ip,omitmpty" gorm:"source_ip,omitmpty"`
		SourcePort string `json:"src_port" gorm:"source_port"`
		UserAgent  string `json:"user_agent" gorm:"user_agent"`
	}
	type DIPUAC struct {
		DestinationIP   string `json:"destination_ip,omitmpty" gorm:"destination_ip,omitmpty"`
		DestinationPort string `json:"dst_port" gorm:"destination_port"`
		UserAgent       string `json:"user_agent" gorm:"user_agent"`
	}
	uacmap := make(map[string]api.TransactionDataUAC)
	var sipuacarray []SIPUAC
	var dipuacarray []DIPUAC
	//将查询到的source_ip和user_agent存入uacarray
	for idx := range tablename {
		var sipuactmp []SIPUAC
		var dipuactmp []DIPUAC
		dbhandletmp := dbhandle.Table(tablename[idx])
		dbhandletmp = dbhandletmp.Where("micro_ts BETWEEN ? AND ?", timefrom*1000, timeto*1000+999)
		if callid != nil {
			dbhandletmp = dbhandletmp.Where("callid = ?", callid[0])
		}
		dbhandletmpsipua := dbhandletmp.Select("source_ip, source_port, user_agent")
		dbhandletmpdipua := dbhandletmp.Select("destination_ip, destination_port, user_agent")
		dbhandletmpsipua.Find(&sipuactmp)
		sipuacarray = append(sipuacarray, sipuactmp...)
		dbhandletmpdipua.Find(&dipuactmp)
		dipuacarray = append(dipuacarray, dipuactmp...)
	}
	//将uacarray中存储的ua信息转换为map形式的uac
	for idx := range sipuacarray {
		var tmp api.TransactionDataUAC
		sipkey := sipuacarray[idx].SourceIp + ":" + sipuacarray[idx].SourcePort
		tmp.Image = ""
		tmp.Agent = sipuacarray[idx].UserAgent
		uacmap[sipkey] = tmp
	}
	for idx := range dipuacarray {
		var tmp api.TransactionDataUAC
		dipkey := dipuacarray[idx].DestinationIP + ":" + dipuacarray[idx].DestinationPort
		tmp.Image = ""
		tmp.Agent = dipuacarray[idx].UserAgent
		uacmap[dipkey] = tmp
	}
	//将uacarray中存储的hosts信息转换为map形式的hosts
	hostsmap := make(map[string]api.TransactionDataHostsContent)

	for idx := range sipuacarray {
		var tmp api.TransactionDataHostsContent
		sipkey := sipuacarray[idx].SourceIp + ":" + sipuacarray[idx].SourcePort
		tmp.IsStp = 0
		hostsmap[sipkey] = tmp
	}

	for idx := range dipuacarray {
		var tmp api.TransactionDataHostsContent
		dipkey := dipuacarray[idx].DestinationIP + ":" + dipuacarray[idx].DestinationPort
		tmp.IsStp = 0
		hostsmap[dipkey] = tmp
	}

	var posidx = 0
	for idx := range hostsmap {
		tmp := hostsmap[idx]
		tmp.Position = posidx
		hostcontenttmp := make(map[string]int)
		hostcontenttmp[idx] = posidx
		tmp.Hosts = append(tmp.Hosts, hostcontenttmp)
		hostsmap[idx] = tmp
		posidx++
	}

	//对calldata进行填充
	for idx := range calldata {
		//填充Protoco
		calldata[idx].Protocol = "sip"
		//填充source_alias
		calldata[idx].SourceAlias = calldata[idx].SourceIp
		//填充destination_alias
		calldata[idx].DestinationAlias = calldata[idx].DestinationIp
		//填充src_id
		calldata[idx].SrcID = calldata[idx].SourceIp + ":" + calldata[idx].SourcePort
		//填充dst_id
		calldata[idx].DstID = calldata[idx].DestinationIp + ":" + calldata[idx].DestinationPort
		//填充method_text
		calldata[idx].MethodText = calldata[idx].Method + " " + calldata[idx].ReplyReason
		//填充destination
		destmp := hostsmap[calldata[idx].SourceIp]
		calldata[idx].Destination = destmp.Position
		//填充callid
		calldata[idx].CallID = callid[0]
		//填充MilliTs
		tmp, err := strconv.Atoi(calldata[idx].MicroTs)
		if err != nil {
			fmt.Printf("Micro_ts = %+v\n", calldata[idx].MicroTs)
			fmt.Printf("tmp = %+v\n", tmp)
		}
		calldata[idx].MilliTs = tmp / 1000
	}

	transdata = api.TransactionData{
		Info: api.TransactionDataInfo{
			CallID:     []api.TransactionDataInfoCallId{},
			TotDur:     "00:00:15",
			StatusCall: 1,
		},
		HostsMap: hostsmap,
		UACMap:   uacmap,
		RTPInfo:  []api.TransactionDataRTPInfo{},
		CallData: calldata,
		Count:    len(calldata),
	}
	return transdata, nil
}

func ExportText(searchtransreq api.SearchTransRequest) (text []api.Text, err error) {
	store := helper.GetMysqlSotrage()
	handle, ok := store.(*mysqls.Store)
	if !ok {
		fmt.Println("get handle not ok")
	}
	dbhandle := handle.Client

	var textdata []api.TextData

	timefrom := searchtransreq.TimeStamp.From
	timeto := searchtransreq.TimeStamp.To
	callid := searchtransreq.Param.Search.CallID

	fromdateutc := utils.TsToDTUTC(timefrom)
	todateutc := utils.TsToDTUTC(timeto)

	// 得到所需查询的数据库表
	var tablename []string
	tablenametmp := ""
	datetmp := fromdateutc
	for i := 0; datetmp.Before(todateutc); {
		tablenametmp = fmt.Sprintf("%s_%s", sipCaputureCallAllTable, utils.DTToDBNSuf(datetmp))
		tablename = append(tablename, tablenametmp)
		if tablenametmp != "" {
			i++
		}

		tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRegAllTable, utils.DTToDBNSuf(datetmp))
		tablename = append(tablename, tablenametmp)
		if tablenametmp != "" {
			i++
		}

		tablenametmp = fmt.Sprintf("%s_%s", sipCaputureRestAllTable, utils.DTToDBNSuf(datetmp))
		tablename = append(tablename, tablenametmp)
		if tablenametmp != "" {
			i++
		}
		datetmp = datetmp.AddDate(0, 0, 1)
	}

	// 对所需数据库表进行查询
	for idx := range tablename {
		var textdatatmp []api.TextData
		dbhandletmp := dbhandle.Table(tablename[idx])
		//		dbhandletmp = dbhandletmp.Where("micro_ts BETWEEN ? AND ?", timefrom*1000, timeto*1000+999)
		for idx := range callid {
			dbhandletmp = dbhandletmp.Where("callid = ?", callid[idx])
		}
		dbhandletmp = dbhandletmp.Order("micro_ts")
		dbhandletmp = dbhandletmp.Select("micro_ts, source_ip, source_port, destination_ip, destination_port,msg")
		dbhandletmp = dbhandletmp.Find(&textdatatmp)

		textdata = append(textdata, textdatatmp...)
	}

	for idx := range textdata {
		var texttmp api.Text
		datetime := time.Unix(0, textdata[idx].MicroTs*1000)
		var datetimes string
		datetimes = fmt.Sprint(datetime)
		texttmp.TextHead = "-----" + datetimes + " " +
			textdata[idx].SourceIp + ":" +
			textdata[idx].SourcePort + " -> " +
			textdata[idx].DestinationIp + ":" +
			textdata[idx].DestinationPort
		texttmp.TextMsg = textdata[idx].Msg
		text = append(text, texttmp)
	}
	return text, nil
}
