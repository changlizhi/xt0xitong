package ml2changliangs

func Bm3ZiDuanZhiDingsJieGou() map[string]interface{} {
	// 字段指定表：主键，字段主键，指定编码，可能值，可能值描述。//字段指定跟字段是一对多的，一个字段有多个指定值
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm3ZiDuanZhiDings,
		SuoYin:     ZhiDingBianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangBIGINT(ZiDuanZhuJian, "20"),
			ZuZhuangVARCHAR(KeNengZhi, "50"),
			ZuZhuangVARCHAR(KeNengZhiMiaoShu, "50"),
			ZuZhuangVARCHAR(ZhiDingBianMa, "50"),
		},
	}
	return shuJu0
}
