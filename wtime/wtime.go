package wtime

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Time
func TimeHour() int                { return time.Now().Hour() }
func TimeMin() int                 { return time.Now().Minute() }
func TimeSec() int                 { return time.Now().Second() }
func DateNow() string              { return time.Now().Format("2006-01-02") }
func TimeNow() string              { return time.Now().Format("15:04:05") }
func TimeNowWithNS() string        { return time.Now().Format("15:04:05.000") }
func DateTime() string             { return time.Now().Format("2006-01-02 15:04:05.000") }
func TimeUTC() string              { return time.Now().String() }
func CurrentUnixTime() int64       { return time.Now().Unix() }
func CurrentUnixNanoTime() int64   { return time.Now().UnixNano() }
func UnitToTime(i int64) time.Time { return time.Unix(i, 0) }
func CurrentMonth() int            { return int(time.Now().Month()) }
func CurrentMonthDay() int         { return time.Now().Day() }

// Date

// PrevDay(time).Format("2006-01-02")
func PrevDay(t time.Time) time.Time { return t.AddDate(0, 0, -1) }
func NextDay(t time.Time) time.Time { return t.AddDate(0, 0, 1) }

func Yesterday() time.Time { return time.Now().AddDate(0, 0, -1) }
func YesterdayStr() string { return Yesterday().Format("2006-01-02") }

func Tomorrow() time.Time { return time.Now().AddDate(0, 0, 1) }
func TomorrowStr() string { return Tomorrow().Format("2006-01-02") }

// func CurrentMonthDate() string   { return DateT(time.Now(), "D") }
func DateUnixTime(year int, month time.Month, day int) int64 {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
}

// check available run time, if beyond current time then EXIT
func CheckTimeToExit(year int, month time.Month, day int) {
	if DateUnixTime(year, month, day) < CurrentUnixTime() {
		fmt.Println("CT System ERROR, Call T0Ny @ tonyw2000@outlook.com")
		os.Exit(-1)
	}
}

// start, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
// end, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")

// in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")
// out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")

// if inTimeSpan(start, end, in) {
//     Println(in, "is between", start, "and", end, ".")
// }

//	if !inTimeSpan(start, end, out) {
//	    Println(out, "is not between", start, "and", end, ".")
//	}
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// Time sleep func
func Sleep(sec int)            { SleepSec(sec) }
func SleepHour(hour int)       { time.Sleep(time.Duration(hour) * time.Hour) }
func SleepMin(min int)         { time.Sleep(time.Duration(min) * time.Minute) }
func SleepSec(sec int)         { time.Sleep(time.Duration(sec) * time.Second) }
func SleepMS(ms int)           { SleepMillisecond(ms) }
func SleepMillisecond(ms int)  { time.Sleep(time.Duration(ms) * time.Millisecond) }
func SleepMicrosecond(mic int) { time.Sleep(time.Duration(mic) * time.Microsecond) }
func SleepNanosecond(nan int)  { time.Sleep(time.Duration(nan) * time.Nanosecond) }

// Format unix time int64 to string
func Date(ti int64, format string) string {
	t := time.Unix(int64(ti), 0)
	return DateT(t, format)
}

// Format unix time string to string
func DateS(ts string, format string) string {
	i, _ := strconv.ParseInt(ts, 10, 64)
	return Date(i, format)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func DateT(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Parse Date use PHP time format. ie: DateParse("2016-01-15", "Y-m-d")
func DateParse(dateString, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, dateString, time.Local)
}

func DaysDifference(strDate string) int {
	// 解析输入日期字符串为时间对象
	layout := "2006-01-02"
	date, err := time.Parse(layout, strDate)
	if err != nil {
		fmt.Println("wrong date format: 2006-01-02", err)
		return 0
	}

	// 获取今天的日期时间对象
	today := time.Now()

	// 计算两个日期之间的天数差
	difference := today.Sub(date).Hours() / 24

	return int(math.Abs(float64(difference)))
}
