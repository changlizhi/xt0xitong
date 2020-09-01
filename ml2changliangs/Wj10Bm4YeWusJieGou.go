package ml2changliangs

func Bm4YeWusJieGou() map[string]interface{} {
	// 业务表：主键，业务编码，业务名称，描述。//这是针对一个业务进行业务流向梳理,业务是必须有方法流的，但方法流必须也是一对多的，所以这里不需要去指定方法流字段
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm4YeWus,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MiaoShu, "150"),
			ZuZhuangVARCHAR(XiTongBianMa, "50"),
		},
	}
	return shuJu0
}
