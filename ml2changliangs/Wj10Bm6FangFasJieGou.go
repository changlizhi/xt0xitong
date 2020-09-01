package ml2changliangs

func Bm6FangFasJieGou() map[string]interface{} {
	// 方法表：主键，方法名，名称（带上名称是因为中文更容易理解），描述。//所有方法流里有的方法都能在这里找到。
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm6FangFas,
		SuoYin:     FangFaMing,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
		},
	}
	return shuJu0
}
