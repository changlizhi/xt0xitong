package ml2changliangs

func Bm1BiaosJieGou() map[string]interface{} {
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm1Biaos,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(ZhuJianBiao, "50"),
			ZuZhuangVARCHAR(XiTongBianMa, "50"),
		},
	}
	return shuJu0
}
