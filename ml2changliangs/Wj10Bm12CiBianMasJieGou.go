package ml2changliangs

func Bm12CiBianMasJieGou() map[string]interface{} {
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm12CiBianMas,
		SuoYin:     FhKongZiFu,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangBIGINT(Bm10Cis+ZhuJian, "20"),
			ZuZhuangBIGINT(Bm11BianMas+ZhuJian, "20"),
		},
	}
	return shuJu0
}
