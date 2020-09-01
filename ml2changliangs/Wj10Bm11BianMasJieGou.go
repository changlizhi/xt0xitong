package ml2changliangs

//一个编码可能是多个词组成的，在描述里体现一下即可。
func Bm11BianMasJieGou() map[string]interface{} {
	shuJu0 := map[string]interface{}{
		CaoZuoKu:   XT0JICHU,
		CaoZuoBiao: Bm11BianMas,
		SuoYin:     BianMa,
		ZhuJian:    ZhuJian,
		ZiDuans: []map[string]interface{}{
			ZuZhuangBIGINT(ZhuJian, "20"),
			ZuZhuangVARCHAR(BianMa, "50"),
			ZuZhuangVARCHAR(MingCheng, "50"),
			ZuZhuangVARCHAR(MiaoShu, "350"),
		},
	}
	return shuJu0
}
