package ml2changliangs

func Bm9FangFaBuZhousJieGou() map[string]interface{} {
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm9FangFaBuZhous,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(MiaoShu, "350"),
			ZuZhuangVARCHAR(GoDaiMa, "2000"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
		},
	}
	return shuJu0
}
