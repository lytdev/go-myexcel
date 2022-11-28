package gtime

import (
	"errors"
	"time"
)

const (
	MYNano      = "2006-01-02 15:04:05.000000000"
	MYMicro     = "2006-01-02 15:04:05.000000"
	MYMil       = "2006-01-02 15:04:05.000"
	MYSec       = "2006-01-02 15:04:05"
	MYCST       = "2006-01-02 15:04:05 +0800 CST"
	MYUTC       = "2006-01-02 15:04:05 +0000 UTC"
	MYDate      = "2006-01-02"
	MYTime      = "15:04:05"
	FBTIME      = "2006-01-02T15:04:05+0800"
	APPTIME     = "2006-01-02T15:04:05.000"
	TWITTERTIME = "2006-01-02T15:04:05Z"
)

var TIME_LOCATION *time.Location

func init() {
	var err error
	TIME_LOCATION, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

// GetCurrentTime 当前时区的当前时间
func GetCurrentTime() time.Time {
	return time.Now().In(TIME_LOCATION)
}

func getTimeDefault() time.Time {
	t, _ := time.ParseInLocation(MYSec, "", TIME_LOCATION)
	return t
}

// Time2StrAsFormat 按照指定的格式输出时间
func Time2StrAsFormat(t time.Time, timeFormat string) string {
	// 先将输入的时间转换到指定的时区，然后再转换格式
	return t.In(TIME_LOCATION).Format(timeFormat)
}

// TimeStr2Time 时间字符串转时间
func TimeStr2Time(timeStr string) (time.Time, error) {
	// 可能的转换格式
	useFormat := []string{
		MYNano, MYMicro, MYMil, MYSec, MYCST, MYUTC, MYDate, MYTime, FBTIME, APPTIME, TWITTERTIME,
		time.RFC3339,
		time.RFC3339Nano,
	}
	var t time.Time
	for _, useF := range useFormat {
		tt, err1 := time.ParseInLocation(useF, timeStr, TIME_LOCATION)
		if err1 != nil {
			continue
		}
		t = tt
		break
	}
	if t == getTimeDefault() { // 0001-01-01 00:00:00 +0000 UTC
		return t, errors.New("时间字符串格式错误")
	}
	return t, nil
}

///////////////////////时间的时区转换//////////////////////

// Time2UTC 本地时区时间与UTC时区时间转换
func Time2UTC(t time.Time) time.Time {
	// 时间转换成 UTC时区的时间
	return t.UTC()
}

// Time2Local 转成本地时区时间
func Time2Local(t time.Time) time.Time {
	return t.Local()
}

//////////////////时间戳与时间的相关转换/////////////////////////

// Time2TimeStampSecond 时间转秒级别时间戳
func Time2TimeStampSecond(t time.Time) int64 {
	return t.Unix()
}

// Time2TimeStampNano 时间转纳秒级别时间戳
func Time2TimeStampNano(t time.Time) int64 {
	return t.UnixNano()
}

// Time2TimeStampMill 时间转毫秒级别时间戳
func Time2TimeStampMill(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// Timestamp2Time 时间戳转时间
func Timestamp2Time(stamp int64, nsec int64) time.Time {
	return time.Unix(stamp, nsec)
}

// TimestampNano2Time 纳秒时间戳转时间
func TimestampNano2Time(stamp int64) time.Time {
	return Timestamp2Time(0, stamp)
}

// TimestampMil2Time 毫秒时间戳转时间(毫秒 *1e6 先转成纳秒)
func TimestampMil2Time(stamp int64) time.Time {
	return Timestamp2Time(0, stamp*1e6)
}

// TimestampSec2Time 秒级别时间戳转时间
func TimestampSec2Time(stamp int64) time.Time {
	return Timestamp2Time(stamp, 0)
}

// TimeStr2TimestampMill 字符串转毫秒时间戳
func TimeStr2TimestampMill(timeStr string) (int64, error) {
	t, err := TimeStr2Time(timeStr)
	if err != nil {
		return -1., err
	}
	// 毫秒级别
	return (t.UnixNano()) / 1e6, nil
}

// ////////////////////////////////////////////////////////////
const TimeActivitiesLayout = "20060102"

// NumberToDate 时间数字int必须是：20210222
func NumberToDate(number int) time.Time {
	var year = number / 10000
	var month = number % 10000 / 100
	var day = number % 100
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, TIME_LOCATION)
}

// StringToDate 时间字符串的格式必须是："20210222"
func StringToDate(s string) (time.Time, error) {
	timeRet, err := time.ParseInLocation(TimeActivitiesLayout, s, TIME_LOCATION)
	if err != nil {
		return timeRet, err
	}
	return timeRet, nil
}