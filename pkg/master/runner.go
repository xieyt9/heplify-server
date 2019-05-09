/*
Copyright (c) 2018 Tinet. All rights reserved.

Author: seanchann <zhouxq@ti-net.com.cn>

See docs/ for more information about this project.
*/

package master

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/config"
	"github.com/sipcapture/heplify-server/pkg/storage/helper"
	"github.com/sipcapture/heplify-server/pkg/storage/mysqls"

	"github.com/jinzhu/now"
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

func createISUPTable(handle *mysqls.Store, tablename string) {
	dbhandle := handle.Client

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		micro_ts bigint(18) NOT NULL DEFAULT '0',
		method varchar(4) NOT NULL DEFAULT '',
		correlation_id varchar(256) NOT NULL DEFAULT '',
		opc int(10) NOT NULL DEFAULT 0,
		dpc int(10) NOT NULL DEFAULT 0,
		cic int(10) NOT NULL DEFAULT 0,
		called_number varchar(16)  DEFAULT '',
		called_ton int(10)  DEFAULT 0,
		called_npi int(10)  DEFAULT 0,
		called_inn int(10)  DEFAULT 0,
		calling_number varchar(16)  DEFAULT '',
		calling_ton int(10)  DEFAULT 0,
		calling_npi int(10)  DEFAULT 0,
		calling_ni int(10)  DEFAULT 0,
		calling_restrict int(10)  DEFAULT 0,
		calling_screened int(10)  DEFAULT 0,
		calling_category int(10)  DEFAULT 0,
		cause_standard int(10)  DEFAULT 0,
		cause_location int(10)  DEFAULT 0,
		cause_itu_class int(10)  DEFAULT 0,
		cause_itu_cause int(10)  DEFAULT 0,
		event_num int(10)  DEFAULT 0,
		hop_counter int(10)  DEFAULT 0,
		nci_satellite int(10)  DEFAULT 0,
		nci_continuity_check int(10)  DEFAULT 0,
		nci_echo_device int(10)  DEFAULT 0,
		fwc_nic int(10)  DEFAULT 0,
		fwc_etem int(10)  DEFAULT 0,
		fwc_iw int(10)  DEFAULT 0,
		fwc_etei int(10)  DEFAULT 0,
		fwc_isup int(10)  DEFAULT 0,
		fwc_isup_pref int(10)  DEFAULT 0,
		fwc_ia int(10)  DEFAULT 0,
		fwc_sccpm int(10)  DEFAULT 0,
		transmission_medium int(10)  DEFAULT 0,
		user_coding_standard int(10)  DEFAULT 0,
		user_transfer_cap int(10)  DEFAULT 0,
		user_transfer_mode int(10)  DEFAULT 0,
		user_transfer_rate int(10)  DEFAULT 0,
		user_layer1_ident int(10)  DEFAULT 0,
		user_layer1_proto int(10)  DEFAULT 0,
		source_ip varchar(60) NOT NULL DEFAULT '',
		source_port int(10) NOT NULL DEFAULT 0,
		destination_ip varchar(60) NOT NULL DEFAULT '',
		destination_port int(10) NOT NULL DEFAULT 0,
		proto int(5) NOT NULL DEFAULT 0,
		family int(1) DEFAULT NULL,
		type int(5) NOT NULL DEFAULT 0,
		node varchar(125) NOT NULL DEFAULT '',
		msg varchar(3000) NOT NULL DEFAULT '',
		expires int(5) NOT NULL DEFAULT '-1',
		PRIMARY KEY (id,date),
		KEY date (date),
		KEY called_number (called_number),
		KEY calling_number (calling_number),
		KEY correlationid (correlation_id(255))
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8 ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8`, tablename)

	dbhandle.Exec(sql)
}

func createRTCPTable(handle *mysqls.Store, tablename string) {
	dbhandle := handle.Client

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
 	    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
 		date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		micro_ts bigint(18) NOT NULL DEFAULT '0',
 		correlation_id varchar(256) NOT NULL DEFAULT '',
 		source_ip varchar(60) NOT NULL DEFAULT '',
 		source_port int(10) NOT NULL DEFAULT 0,
 		destination_ip varchar(60) NOT NULL DEFAULT '',
 		destination_port int(10) NOT NULL DEFAULT 0,
 		proto int(5) NOT NULL DEFAULT 0,
 		family int(1) DEFAULT NULL,
 		type int(5) NOT NULL DEFAULT 0,
 		node varchar(125) NOT NULL DEFAULT '',
 		msg varchar(1500) NOT NULL DEFAULT '',
 		PRIMARY KEY (id,date),
 		KEY date (date),
 		KEY correlationid (correlation_id(255))
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8 ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8`, tablename)

	dbhandle.Exec(sql)
}

func createSIPCallTable(handle *mysqls.Store, tablename string) {
	dbhandle := handle.Client
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
 		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
 		date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		micro_ts bigint(18) NOT NULL DEFAULT '0',
 		method varchar(50) NOT NULL DEFAULT '',
 		reply_reason varchar(100) NOT NULL DEFAULT '',
 		ruri varchar(200) NOT NULL DEFAULT '',
 		ruri_user varchar(100) NOT NULL DEFAULT '',
 		ruri_domain varchar(150) NOT NULL DEFAULT '',
 		from_user varchar(100) NOT NULL DEFAULT '',
 		from_domain varchar(150) NOT NULL DEFAULT '',
 		from_tag varchar(64) NOT NULL DEFAULT '',
 		to_user varchar(100) NOT NULL DEFAULT '',
 		to_domain varchar(150) NOT NULL DEFAULT '',
 		to_tag varchar(64) NOT NULL DEFAULT '',
 		pid_user varchar(100) NOT NULL DEFAULT '',
 		contact_user varchar(120) NOT NULL DEFAULT '',
 		auth_user varchar(120) NOT NULL DEFAULT '',
 		callid varchar(120) NOT NULL DEFAULT '',
 		callid_aleg varchar(120) NOT NULL DEFAULT '',
 		via_1 varchar(256) NOT NULL DEFAULT '',
 		via_1_branch varchar(80) NOT NULL DEFAULT '',
 		cseq varchar(25) NOT NULL DEFAULT '',
 		diversion varchar(256) NOT NULL DEFAULT '',
 		reason varchar(200) NOT NULL DEFAULT '',
 		content_type varchar(256) NOT NULL DEFAULT '',
 		auth varchar(256) NOT NULL DEFAULT '',
 		user_agent varchar(256) NOT NULL DEFAULT '',
 		source_ip varchar(60) NOT NULL DEFAULT '',
 		source_port int(10) NOT NULL DEFAULT 0,
 		destination_ip varchar(60) NOT NULL DEFAULT '',
 		destination_port int(10) NOT NULL DEFAULT 0,
 		contact_ip varchar(60) NOT NULL DEFAULT '',
 		contact_port int(10) NOT NULL DEFAULT 0,
 		originator_ip varchar(60) NOT NULL DEFAULT '',
 		originator_port int(10) NOT NULL DEFAULT 0,
 		expires int(5) NOT NULL DEFAULT '-1',
 		correlation_id varchar(256) NOT NULL DEFAULT '',
 		custom_field1 varchar(120) NOT NULL DEFAULT '',
 		custom_field2 varchar(120) NOT NULL DEFAULT '',
 		custom_field3 varchar(120) NOT NULL DEFAULT '',
 		proto int(5) NOT NULL DEFAULT 0,
 		family int(1) DEFAULT NULL,
 		rtp_stat varchar(256) NOT NULL DEFAULT '',
 		type int(2) NOT NULL DEFAULT 0,
 		node varchar(125) NOT NULL DEFAULT '',
 		msg varchar(1500) NOT NULL DEFAULT '',
		PRIMARY KEY (id,date),
		KEY from_user (from_user),
		KEY to_user (to_user),
 		KEY date (date),
 		KEY callid (callid)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8`, tablename)

	dbhandle.Exec(sql)
}
func createSIPRegTable(handle *mysqls.Store, tablename string) {
	dbhandle := handle.Client
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
 		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
 		date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		micro_ts bigint(18) NOT NULL DEFAULT '0',
 		method varchar(50) NOT NULL DEFAULT '',
 		reply_reason varchar(100) NOT NULL DEFAULT '',
 		ruri varchar(200) NOT NULL DEFAULT '',
 		ruri_user varchar(100) NOT NULL DEFAULT '',
 		ruri_domain varchar(150) NOT NULL DEFAULT '',
 		from_user varchar(100) NOT NULL DEFAULT '',
 		from_domain varchar(150) NOT NULL DEFAULT '',
 		from_tag varchar(64) NOT NULL DEFAULT '',
 		to_user varchar(100) NOT NULL DEFAULT '',
 		to_domain varchar(150) NOT NULL DEFAULT '',
 		to_tag varchar(64) NOT NULL DEFAULT '',
 		pid_user varchar(100) NOT NULL DEFAULT '',
 		contact_user varchar(120) NOT NULL DEFAULT '',
 		auth_user varchar(120) NOT NULL DEFAULT '',
 		callid varchar(120) NOT NULL DEFAULT '',
 		callid_aleg varchar(120) NOT NULL DEFAULT '',
 		via_1 varchar(256) NOT NULL DEFAULT '',
 		via_1_branch varchar(80) NOT NULL DEFAULT '',
 		cseq varchar(25) NOT NULL DEFAULT '',
 		diversion varchar(256) NOT NULL DEFAULT '',
 		reason varchar(200) NOT NULL DEFAULT '',
 		content_type varchar(256) NOT NULL DEFAULT '',
 		auth varchar(256) NOT NULL DEFAULT '',
 		user_agent varchar(256) NOT NULL DEFAULT '',
 		source_ip varchar(60) NOT NULL DEFAULT '',
 		source_port int(10) NOT NULL DEFAULT 0,
 		destination_ip varchar(60) NOT NULL DEFAULT '',
 		destination_port int(10) NOT NULL DEFAULT 0,
 		contact_ip varchar(60) NOT NULL DEFAULT '',
 		contact_port int(10) NOT NULL DEFAULT 0,
 		originator_ip varchar(60) NOT NULL DEFAULT '',
 		originator_port int(10) NOT NULL DEFAULT 0,
 		expires int(5) NOT NULL DEFAULT '-1',
 		correlation_id varchar(256) NOT NULL DEFAULT '',
 		custom_field1 varchar(120) NOT NULL DEFAULT '',
 		custom_field2 varchar(120) NOT NULL DEFAULT '',
 		custom_field3 varchar(120) NOT NULL DEFAULT '',
 		proto int(5) NOT NULL DEFAULT 0,
 		family int(1) DEFAULT NULL,
 		rtp_stat varchar(256) NOT NULL DEFAULT '',
 		type int(2) NOT NULL DEFAULT 0,
 		node varchar(125) NOT NULL DEFAULT '',
 		msg varchar(1500) NOT NULL DEFAULT '',
 		PRIMARY KEY (id,date),
 		KEY from_user (from_user),
		 KEY to_user (to_user),
		 KEY callid (callid),
 		KEY date (date)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8`, tablename)

	dbhandle.Exec(sql)
}

func createSIPRestTable(handle *mysqls.Store, tablename string) {
	dbhandle := handle.Client
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
 		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
 		date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		micro_ts bigint(18) NOT NULL DEFAULT '0',
 		method varchar(50) NOT NULL DEFAULT '',
 		reply_reason varchar(100) NOT NULL DEFAULT '',
 		ruri varchar(200) NOT NULL DEFAULT '',
 		ruri_user varchar(100) NOT NULL DEFAULT '',
 		ruri_domain varchar(150) NOT NULL DEFAULT '',
 		from_user varchar(100) NOT NULL DEFAULT '',
 		from_domain varchar(150) NOT NULL DEFAULT '',
 		from_tag varchar(64) NOT NULL DEFAULT '',
 		to_user varchar(100) NOT NULL DEFAULT '',
 		to_domain varchar(150) NOT NULL DEFAULT '',
 		to_tag varchar(64) NOT NULL DEFAULT '',
 		pid_user varchar(100) NOT NULL DEFAULT '',
 		contact_user varchar(120) NOT NULL DEFAULT '',
 		auth_user varchar(120) NOT NULL DEFAULT '',
 		callid varchar(120) NOT NULL DEFAULT '',
 		callid_aleg varchar(120) NOT NULL DEFAULT '',
 		via_1 varchar(256) NOT NULL DEFAULT '',
 		via_1_branch varchar(80) NOT NULL DEFAULT '',
 		cseq varchar(25) NOT NULL DEFAULT '',
 		diversion varchar(256) NOT NULL DEFAULT '',
 		reason varchar(200) NOT NULL DEFAULT '',
 		content_type varchar(256) NOT NULL DEFAULT '',
 		auth varchar(256) NOT NULL DEFAULT '',
 		user_agent varchar(256) NOT NULL DEFAULT '',
 		source_ip varchar(60) NOT NULL DEFAULT '',
 		source_port int(10) NOT NULL DEFAULT 0,
 		destination_ip varchar(60) NOT NULL DEFAULT '',
 		destination_port int(10) NOT NULL DEFAULT 0,
 		contact_ip varchar(60) NOT NULL DEFAULT '',
 		contact_port int(10) NOT NULL DEFAULT 0,
 		originator_ip varchar(60) NOT NULL DEFAULT '',
 		originator_port int(10) NOT NULL DEFAULT 0,
 		expires int(5) NOT NULL DEFAULT '-1',
 		correlation_id varchar(256) NOT NULL DEFAULT '',
 		custom_field1 varchar(120) NOT NULL DEFAULT '',
 		custom_field2 varchar(120) NOT NULL DEFAULT '',
 		custom_field3 varchar(120) NOT NULL DEFAULT '',
 		proto int(5) NOT NULL DEFAULT 0,
 		family int(1) DEFAULT NULL,
 		rtp_stat varchar(256) NOT NULL DEFAULT '',
 		type int(2) NOT NULL DEFAULT 0,
 		node varchar(125) NOT NULL DEFAULT '',
 		msg varchar(1500) NOT NULL DEFAULT '',
 		PRIMARY KEY (id,date),
 		KEY ruri_user (ruri_user),
 		KEY from_user (from_user),
 		KEY to_user (to_user),
 		KEY pid_user (pid_user),
 		KEY auth_user (auth_user),
 		KEY callid_aleg (callid_aleg),
 		KEY date (date),
 		KEY callid (callid)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8`, tablename)

	dbhandle.Exec(sql)
}

func dbMaintain() error {
	store := helper.GetMysqlSotrage()
	handle, ok := store.(*mysqls.Store)
	if !ok {
		return fmt.Errorf("need mysql handle")
	}

	dbhandle := handle.Client

	//use CTS +8 timezone with this work.
	loc, _ := time.LoadLocation("Asia/Shanghai")
	nowtime := time.Now().In(loc)

	dropTableDuration := time.Duration(time.Hour * 24 * time.Duration(config.Setting.DropTableDays))
	newTableDuration := time.Duration(time.Hour * 24)
	nextTableDuration := time.Duration(time.Hour * 24 * 2)

	dropTableDate := nowtime.Add(-dropTableDuration)
	newTableDate := nowtime.Add(newTableDuration)
	nextTableDate := nowtime.Add(nextTableDuration)

	dropTableISUP := fmt.Sprintf("%s_%d%02d%02d", isupCaputureAllTable, dropTableDate.Year(), int(dropTableDate.Month()), dropTableDate.Day())
	dropTableRTCP := fmt.Sprintf("%s_%d%02d%02d", rtcpCaputureAllTable, dropTableDate.Year(), int(dropTableDate.Month()), dropTableDate.Day())
	dropTableSIPCall := fmt.Sprintf("%s_%d%02d%02d", sipCaputureCallAllTable, dropTableDate.Year(), int(dropTableDate.Month()), dropTableDate.Day())
	dropTableSIPReg := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRegAllTable, dropTableDate.Year(), int(dropTableDate.Month()), dropTableDate.Day())
	dropTableSIPRest := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRestAllTable, dropTableDate.Year(), int(dropTableDate.Month()), dropTableDate.Day())

	glog.Infof("database DropTableDays(%d) dropTableDate(%v) dropTableISUP(%s)", config.Setting.DropTableDays, dropTableDate, dropTableISUP)
	dbhandle.DropTableIfExists(dropTableISUP)
	dbhandle.DropTableIfExists(dropTableRTCP)
	dbhandle.DropTableIfExists(dropTableSIPCall)
	dbhandle.DropTableIfExists(dropTableSIPReg)
	dbhandle.DropTableIfExists(dropTableSIPRest)

	TodayTableISUP := fmt.Sprintf("%s_%d%02d%02d", isupCaputureAllTable, nowtime.Year(), int(nowtime.Month()), nowtime.Day())
	TodayTableRTCP := fmt.Sprintf("%s_%d%02d%02d", rtcpCaputureAllTable, nowtime.Year(), int(nowtime.Month()), nowtime.Day())
	TodayTableSIPCall := fmt.Sprintf("%s_%d%02d%02d", sipCaputureCallAllTable, nowtime.Year(), int(nowtime.Month()), nowtime.Day())
	TodayTableSIPReg := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRegAllTable, nowtime.Year(), int(nowtime.Month()), nowtime.Day())
	TodayTableSIPRest := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRestAllTable, nowtime.Year(), int(nowtime.Month()), nowtime.Day())

	createISUPTable(handle, TodayTableISUP)
	createRTCPTable(handle, TodayTableRTCP)
	createSIPCallTable(handle, TodayTableSIPCall)
	createSIPRegTable(handle, TodayTableSIPReg)
	createSIPRestTable(handle, TodayTableSIPRest)

	newTableISUP := fmt.Sprintf("%s_%d%02d%02d", isupCaputureAllTable, newTableDate.Year(), int(newTableDate.Month()), newTableDate.Day())
	newTableRTCP := fmt.Sprintf("%s_%d%02d%02d", rtcpCaputureAllTable, newTableDate.Year(), int(newTableDate.Month()), newTableDate.Day())
	newTableSIPCall := fmt.Sprintf("%s_%d%02d%02d", sipCaputureCallAllTable, newTableDate.Year(), int(newTableDate.Month()), newTableDate.Day())
	newTableSIPReg := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRegAllTable, newTableDate.Year(), int(newTableDate.Month()), newTableDate.Day())
	newTableSIPRest := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRestAllTable, newTableDate.Year(), int(newTableDate.Month()), newTableDate.Day())

	createISUPTable(handle, newTableISUP)
	createRTCPTable(handle, newTableRTCP)
	createSIPCallTable(handle, newTableSIPCall)
	createSIPRegTable(handle, newTableSIPReg)
	createSIPRestTable(handle, newTableSIPRest)

	nextTableISUP := fmt.Sprintf("%s_%d%02d%02d", isupCaputureAllTable, nextTableDate.Year(), int(nextTableDate.Month()), nextTableDate.Day())
	nextTableRTCP := fmt.Sprintf("%s_%d%02d%02d", rtcpCaputureAllTable, nextTableDate.Year(), int(nextTableDate.Month()), nextTableDate.Day())
	nextTableSIPCall := fmt.Sprintf("%s_%d%02d%02d", sipCaputureCallAllTable, nextTableDate.Year(), int(nextTableDate.Month()), nextTableDate.Day())
	nextTableSIPReg := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRegAllTable, nextTableDate.Year(), int(nextTableDate.Month()), nextTableDate.Day())
	nextTableSIPRest := fmt.Sprintf("%s_%d%02d%02d", sipCaputureRestAllTable, nextTableDate.Year(), int(nextTableDate.Month()), nextTableDate.Day())

	createISUPTable(handle, nextTableISUP)
	createRTCPTable(handle, nextTableRTCP)
	createSIPCallTable(handle, nextTableSIPCall)
	createSIPRegTable(handle, nextTableSIPReg)
	createSIPRestTable(handle, nextTableSIPRest)

	return nil
}

//DatabaseMaintainWorker ensure the new table name with date create and
//delete old table
func DatabaseMaintainWorker(wg *sync.WaitGroup, quit chan struct{}) {
	var innerWg sync.WaitGroup

	wg.Add(1)
	go func() {
		//use CTS +8 timezone with this work.
		loc, _ := time.LoadLocation("Asia/Shanghai")
		nowtime := time.Now().In(loc)
		endDate := now.EndOfDay()

		workTime := endDate.Sub(nowtime)
		glog.Infof("database maintain now(%v) endday(%v) worktime(%v)", nowtime, endDate, workTime.String())
		for {
			select {
			case <-time.After(workTime):
				innerWg.Add(1)

				nowtime = time.Now().In(loc)
				endDate = now.EndOfDay()

				workTime = endDate.Sub(nowtime)
				dbMaintain()
				innerWg.Done()
				glog.Infof("database maintain now(%v) endday(%v) worktime(%v)", nowtime, endDate, workTime.String())
			case <-quit:
				innerWg.Wait()
				wg.Done()
				return
			}
		}
	}()
}
