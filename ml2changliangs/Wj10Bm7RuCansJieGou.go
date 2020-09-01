package ml2changliangs

func Bm7RuCansJieGou() map[string]interface{} {
	//入参字段表：主键，编码，名称，方法名（不用方法主键是为了清晰），类型，是否必须，父编码，描述。
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm7RuCans,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(FangFaMing, "50"),
			ZuZhuangVARCHAR(LeiXing, "50"),
			ZuZhuangINT(BiXu, "1"),
			ZuZhuangVARCHAR(FuBianMa, "50"),
			ZuZhuangVARCHAR(MiaoShu, "50"),
		},
	}
	return shuJu0
}
