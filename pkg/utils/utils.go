package utils

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type JsonTime time.Time

// json序列化，以适应前端时间格式
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Unxi时间转换为数据库表中时间
func TsToDT(unixtime int64) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	datetime := time.Unix(unixtime/1000-3600*8, 0).In(loc)
	return datetime
}

// Unix时间转为datetime
func TsToDTUTC(unixtime int64) time.Time {
	datetime := time.Unix(unixtime/1000, 0)
	return datetime
}

// 本地日期转换为数据库表名后缀
func DTToDBNSuf(datetime time.Time) string {
	year := datetime.Year()
	month := datetime.Month()
	day := datetime.Day()

	suffix := fmt.Sprintf("%d%02d%02d", year, month, day)
	return suffix
}

// SplitParam解析输入的参数，参数以'/'分隔
// splitcode:
// 1:str为user
// 2:str为ip
// 3:str为user/ip
// 0:str为其它
func SplitParam(str string) (splitcode int, splitresult []string) {
	splitresult = strings.Split(str, "/")
	if len(splitresult) == 0 {
		return 0, nil
	} else if len(splitresult) == 1 && splitresult[0] != "" {
		if net.ParseIP(splitresult[0]) == nil {
			return 1, splitresult
		} else {
			return 2, splitresult
		}
	} else if len(splitresult) == 2 {
		if net.ParseIP(splitresult[0]) == nil && net.ParseIP(splitresult[1]) != nil {
			return 3, splitresult
		} else {
			return 0, nil
		}
	} else {
		return 0, nil
	}
}
