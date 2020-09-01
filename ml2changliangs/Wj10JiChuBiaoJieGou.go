package ml2changliangs

func HuoQuJiChuBiaoJieGou(biaoMing string) map[string]interface{} {
	ret := map[string]interface{}{
		Bm0XiTongs:        Bm0XiTongsJieGou(),
		Bm1Biaos:          Bm1BiaosJieGou(),
		Bm2ZiDuans:        Bm2ZiDuansJieGou(),
		Bm3ZiDuanZhiDings: Bm3ZiDuanZhiDingsJieGou(),
		Bm4YeWus:          Bm4YeWusJieGou(),
		Bm5YeWuFangFaLius: Bm5YeWuFangFaLiusJieGou(),
		Bm6FangFas:        Bm6FangFasJieGou(),
		Bm7RuCans:         Bm7RuCansJieGou(),
		Bm8ChuCans:        Bm8ChuCansJieGou(),
		Bm9FangFaBuZhous:  Bm9FangFaBuZhousJieGou(),
		Bm10Cis:           Bm10CisJieGou(),
		Bm11BianMas:       Bm11BianMasJieGou(),
		Bm12CiBianMas:     Bm12CiBianMasJieGou(),
	}
	if biaoMing == FhKongZiFu {
		return ret
	}
	return ret[biaoMing].(map[string]interface{})
}
