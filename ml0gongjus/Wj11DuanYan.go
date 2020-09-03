package ml0gongjus

import (
	"xt0xitong/ml2changliangs"
)

func DuanYanZiFuChuan(canShu interface{}) string {
	if canShu != nil {
		return canShu.(string)
	}
	return ml2changliangs.FhKongZiFu
}
