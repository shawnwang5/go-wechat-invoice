package utils

import (
	"fmt"
	"time"
)

// 将毫秒日期字符串转换成日期对象
func MillisecondDatetimeStrToDate(datetimeStr string) (*time.Time, error) {
	timeLayout := "2006-01-02 15:04:05.999"
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}
	datetime, err := time.ParseInLocation(timeLayout, datetimeStr, loc)
	if err != nil {
		return nil, err
	}
	datetime = datetime.UTC()
	return &datetime, nil
}

// 将日期对象转换成毫秒日期字符串
func DateToMillisecondDatetimeStr(time time.Time) string {
	layout := "2006-01-02 15:04:05.999"
	return time.Format(layout)
}

// 将秒日期字符串转换成日期对象
func SecondDatetimeStrToDate(datetimeStr string) (*time.Time, error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}
	datetime, err := time.ParseInLocation(timeLayout, datetimeStr, loc)
	if err != nil {
		return nil, err
	}
	datetime = datetime.UTC()
	return &datetime, nil
}

// 将日期对象转换成秒日期字符串
func DateToSecondDatetimeStr(time time.Time) string {
	layout := "2006-01-02 15:04:05"
	return time.Format(layout)
}

// AddDays 给定的日期字符串加days天， dataStr like [2024-12-17]
func AddDays(dateStr string, days int) (string, error) {
	// 解析日期字符串
	t, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return "", err
	}

	// 在t日期上加 days天
	nextDay := t.AddDate(0, 0, days)

	// 返回格式化后的新日期字符串
	return nextDay.Format(time.DateOnly), nil
}

// AddDays 给定的日期字符串加days天， dataStr like [2024-12-17]
func AddDays2Time(t time.Time, days int) (time.Time, error) {

	// 在t日期上加 days天
	nextDay := t.AddDate(0, 0, days)

	// 返回格式化后的新日期字符串
	return nextDay, nil
}

// 返回整天
func GetDateAfterNow(days int) (*time.Time, error) {
	t1 := time.Now().AddDate(0, 0, days)
	fmt.Println(t1.Format(time.DateTime))

	dateStr := t1.Format("2006-01-02 00:00:00")
	fmt.Println(dateStr)

	// 解析日期字符串
	t, err2 := time.Parse(time.DateTime, dateStr)
	if err2 != nil {
		return nil, err2
	}
	return &t, nil
}

// AddMinutes 给定的日期字符串加minutes分， dataStr like [2024-12-17 09:45:33]
func AddMinutes(dateStr string, minutes int) (string, error) {
	// 解析日期字符串
	t, err := time.Parse(time.DateTime, dateStr)
	if err != nil {
		return "", err
	}

	// 在t日期上加 minutes分钟
	nextDay := t.Add(time.Duration(minutes * int(time.Minute)))

	// 返回格式化后的新日期字符串
	return nextDay.Format(time.DateTime), nil
}

// AddMinutes2 给定的日期字符串加minutes分， dataStr like [2024-12-17 09:45]
func AddMinutes2(dateStr string, minutes int) (string, error) {
	MinuteTime := "2006-01-02 15:04"

	// 解析日期字符串
	t, err := time.Parse(MinuteTime, dateStr)
	if err != nil {
		return "", err
	}

	// 在t日期上加 minutes分钟
	nextDay := t.Add(time.Duration(minutes * int(time.Minute)))

	// 返回格式化后的新日期字符串
	return nextDay.Format(MinuteTime), nil
}
