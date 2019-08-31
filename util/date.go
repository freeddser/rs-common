package util

import (
	"fmt"
	"time"

	"github.com/freeddser/rs-common/config"
)

var (
	timezoneLoc *time.Location
)

func InitTimeZoneLocation() {
	serverMode := config.MustGetString("server.mode")
	timezoneLocString := config.MustGetString(serverMode + ".timezone")
	SetTimeZoneLocation(timezoneLocString)
}

func GetTimeZoneLocation() *time.Location {
	return timezoneLoc
}

func SetTimeZoneLocation(timezoneLocString string) {
	timezoneLoc, _ = time.LoadLocation(timezoneLocString)
}

func IsNowDate(date string) bool {
	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		log.Error(err)
	}

	timeOne := t.Truncate(24 * time.Hour)
	timeTwo := GetTimeNow().Truncate(24 * time.Hour)

	return timeOne.Equal(timeTwo)
}

func IsSameDate(date string, dateTwo string) bool {
	t, err := time.Parse(time.RFC3339, date)
	t2, err := time.Parse(time.RFC3339, dateTwo)

	if err != nil {
		log.Error(err)
	}

	timeOne := t.Truncate(24 * time.Hour)
	timeTwo := t2.Truncate(24 * time.Hour)

	return timeOne.Equal(timeTwo)
}

func IsSameTime(date string, dateTwo string) bool {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", date, timezoneLoc)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", dateTwo, timezoneLoc)

	if err != nil {
		log.Error(err)
	}

	timeOne := t.Truncate(24 * time.Hour)
	timeTwo := t2.Truncate(24 * time.Hour)

	return timeOne.Equal(timeTwo)
}

func ToTimeRFC3339(unixTime int64) string {
	unixTimeUTC := time.Unix(unixTime, 0)
	return unixTimeUTC.Format(time.RFC3339)
}

func ToEpoch(date string) int64 {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Error(err)
	}
	return t.Unix()
}

func ToEpochWIthTruncate(date string) int64 {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Error(err)
	}
	timeTruncate := t.Truncate(24 * time.Hour)
	return timeTruncate.Unix()
}

func ToTimeDate(date string) (resDate string, resTime string) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Error(err)
	}

	hour, min, _ := t.Clock()

	year, month, day := t.Date()

	resTm := fmt.Sprintf("%02d:%02d", hour, min)
	resDt := fmt.Sprintf("%02d-%02d-%d", day, month, year)

	return resDt, resTm
}

func ConvertDateAndTimeStringToTime(dateStr string, timeStr string) time.Time {
	dateTime := dateStr + " " + timeStr + ""
	//log.Info(dateTime)
	t, err := time.ParseInLocation("02-01-2006 15:04:05", dateTime, timezoneLoc)
	//log.Info(t)
	if err != nil {
		log.Error(err)
	}
	return t
}

func GetTimeNow() time.Time {
	//init the loc
	//set timezone,
	now := time.Now().In(timezoneLoc)

	return now
}

func AddTime(dateTime time.Time, d time.Duration) time.Time {
	return dateTime.Add(d)
}

func ToTimeString(time time.Time) (resDate string) {
	return time.In(timezoneLoc).Format("2006-01-02 15:04:05")
}

func ToTimeRFC3339String(date time.Time) (resDate string) {
	return date.Format(time.RFC3339)
}

func GetBeginningYesterday() (timeYesterday time.Time) {
	timeStrBeginningYesterday := GetTimeNow().Add(-24 * time.Hour).Format("2006-01-02 00:00:00")
	t, err := time.ParseInLocation("2006-01-02 00:00:00", timeStrBeginningYesterday, time.Local)
	if err != nil {
		log.Error(err)
	}
	return t
}

func GetBeginningToday() (timeYesterday time.Time) {

	timeStrBeginningYesterday := GetTimeNow().Format("2006-01-02 00:00:00")
	t, err := time.ParseInLocation("2006-01-02 00:00:00", timeStrBeginningYesterday, time.Local)
	if err != nil {
		log.Error(err)
	}
	return t
}

func GetBeginningTomorrow() (timeTomorrow time.Time) {
	timeStrBeginningTomorrow := GetTimeNow().Add(24 * time.Hour).Format("2006-01-02 00:00:00")
	t, err := time.Parse("2006-01-02 00:00:00", timeStrBeginningTomorrow)
	if err != nil {
		log.Error(err)
	}
	return t
}

// todayStart returns today epoch that starts from 00:00:00
func TodayStart() int64 {
	return time.Now().Truncate(24 * time.Hour).Unix()
}

// RelativeEpoch returns now epoch + hours of local time difference
func RelativeEpoch(diff int8) int64 {
	return time.Now().Unix() + (int64(diff) * 3600)
}

func LastDayTime(Number int) int64 {
	return time.Now().Truncate(24*time.Hour).AddDate(0, 0, Number).Unix()
}

func ChangLocalTimeToUTC(localTime time.Time) time.Time {
	utcLocal, err := time.LoadLocation("")
	if err != nil {
		log.Error(err)
		return localTime
	}
	utcTime := localTime.In(utcLocal)
	return utcTime
}

func StringToTime(t string) time.Time {
	var timeLayoutStr = "2006-01-02 15:04:05"
	st, _ := time.Parse(timeLayoutStr, t) //string转time
	return st
}
