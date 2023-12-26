package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Md5 加密
func Md5(str string) string {
	//data := []byte(str)
	//return fmt.Sprintf("%x\n", md5.Sum(data))

	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))

}

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// DateToUnix 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetUnix 获取时间戳秒
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetUnixNano 获取时间戳纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// GetDate 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// GetDay 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// StringToInt 将string转换成int
func StringToInt(str string) (int, error) {
	int, err := strconv.Atoi(str)
	return int, err
}

// IntToString 将int转换成string
func IntToString(int int) string {
	str := strconv.Itoa(int)
	return str
}

// StringToInt64 将string转换成int64
func StringToInt64(str string) (int64, error) {
	int, err := strconv.ParseInt(str, 10, 64)
	return int, err
}

// Int64ToString 将int64转换成string
func Int64ToString(int64 int64) string {
	str := strconv.FormatInt(int64, 10)
	return str
}

// Float64 将string转换成float64
func Float64(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}
