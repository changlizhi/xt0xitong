package ml0gongjus

import (
	"time"
	"xt0xitong/ml2changliangs"
)

func String2Time(str, tmpl string) time.Time {
	//获取本地location
	loc, _ := time.LoadLocation(ml2changliangs.Local)  //重要：获取时区
	theTime, _ := time.ParseInLocation(tmpl, str, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

//str必须为 yyyy-MM-dd hh:mm:ss
func String2TimeNyrSfm(str string) time.Time {
	return String2Time(str, ml2changliangs.NyrSfm)
}

//str必须为 yyyy-MM-dd
func String2TimeNyr(str string) time.Time {
	return String2Time(str, ml2changliangs.Nyr)
}
func Time2String(t time.Time, templ string) string {
	sr := t.Unix() //转化为时间戳 类型是int64
	//时间戳转日期
	dataTimeStr := time.Unix(sr, ml2changliangs.Sz0).Format(templ) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}
func Time2StringNyr(t time.Time) string {
	ret := Time2String(t, ml2changliangs.Nyr)
	return ret
}
func Time2StringNyrSfm(t time.Time) string {
	ret := Time2String(t, ml2changliangs.NyrSfm)
	return ret
}
func Time2StringNyrSfmXhx(t time.Time) string {
	ret := Time2String(t, ml2changliangs.NyrSfmXhx)
	return ret
}
func Time2StringNyrXhx(t time.Time) string {
	ret := Time2String(t, ml2changliangs.NyrXhx)
	return ret
}
func Time2StringNyrWu(t time.Time) string {
	ret := Time2String(t, ml2changliangs.NyrWu)
	return ret
}
