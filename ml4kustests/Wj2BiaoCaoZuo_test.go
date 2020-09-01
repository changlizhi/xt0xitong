package ml4kustests

import (
	"testing"
	"xt0xitong/ml2changliangs"
	"xt0xitong/ml3kus"
)

func TestChuangJianBiaoMingBiao(t *testing.T) {
	shuJu0 := ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{})
	ml3kus.ChuangJianBiao(shuJu0)
}

func TestSheZhiWeiYiSuoYin(t *testing.T) {
	shuJu0 := ml2changliangs.JiChuBiao[ml2changliangs.Bm1BiaoMings].(map[string]interface{})
	ml3kus.SheZhiWeiYiSuoYin(shuJu0)
}
