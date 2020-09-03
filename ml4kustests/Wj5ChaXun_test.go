package ml4kustests

import (
	"testing"
	"xt0xitong/ml2changliangs"
	"xt0xitong/ml3kus"
)

func TestChaXun(t *testing.T) {
	shuJu0 := ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.Bm1BiaoMings)

	bm1BiaoMingsZiDuans := shuJu0[ml2changliangs.ZiDuans].([]map[string]interface{})

	ml3kus.ChuangJianJiChuBiao()
	//初始化一些数据进去
	ml3kus.XinZengYeWuJieGous()

	chaXunZiDuans := []interface{}{}

	for _, v := range bm1BiaoMingsZiDuans {
		chaXunZiDuans = append(chaXunZiDuans, v[ml2changliangs.ZiDuanMing])
	}

	shuJu1 := map[string]interface{}{
		ml2changliangs.CaoZuoKu:   ml2changliangs.XT0JICHU,
		ml2changliangs.CaoZuoBiao: ml2changliangs.Bm1BiaoMings,
		ml2changliangs.ZiDuans:    chaXunZiDuans,
		ml2changliangs.TiaoJianHeZhis: map[string]interface{}{
			ml2changliangs.BianMa: ml2changliangs.Ywb1YongHus, //查询这个条件的数据
		},
	}
	ml3kus.ChaXun(shuJu1)
}
