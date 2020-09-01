package ml0gongjus

import (
	"log"
	"strconv"
	"xt0xitong/ml2changliangs"
)

func String2Float64(str string) float64 {
	ret, err := strconv.ParseFloat(str, ml2changliangs.JinZhi64)
	if err != nil {
		return 0
	}
	return ret
}
func Int642String(shuzi int64) string {
	ret := strconv.FormatInt(shuzi, ml2changliangs.JinZhi10)
	return ret
}
func String2Int64(shuzi string) int64 {
	ret, err := strconv.ParseInt(shuzi, ml2changliangs.JinZhi10, ml2changliangs.JinZhi64)
	if err != nil {
		log.Println(shuzi, err.Error())
	}
	return ret
}
func String2Int(shuzi string) int {
	ret, err := strconv.Atoi(shuzi)
	if err != nil {
		log.Println(shuzi, err.Error())
	}
	return ret
}
func Float642String(shuzi float64) string {
	ret := strconv.FormatFloat(shuzi, 'f', -1, ml2changliangs.JinZhi64)
	return ret
}
