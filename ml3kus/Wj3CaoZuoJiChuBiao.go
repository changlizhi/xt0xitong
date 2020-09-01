package ml3kus

import (
	"xt0xitong/ml2changliangs"
)

func ChuangJianJiChuBiao() {
	for _, v := range ml2changliangs.HuoQuJiChuBiaoJieGou(ml2changliangs.FhKongZiFu) {
		ChuangJianBiao(v.(map[string]interface{}))
	}
}
