package ml2changliangs

func ZuZhuangBIGINT(mingCheng, kuandu string) map[string]interface{} {
	ret := map[string]interface{}{
		ZiDuanMing: mingCheng,
		LeiXing:    BIGINT,
		ChangDu:    kuandu,
		MoRenZhi:   Zf0,
	}
	return ret
}

func ZuZhuangINT(mingCheng, kuandu string) map[string]interface{} {
	ret := map[string]interface{}{
		ZiDuanMing: mingCheng,
		LeiXing:    INT,
		ChangDu:    kuandu,
		MoRenZhi:   Zf0,
	}
	return ret
}

func ZuZhuangVARCHAR(mingCheng, ziFuShu string) map[string]interface{} {
	ret := map[string]interface{}{
		ZiDuanMing: mingCheng,
		LeiXing:    VARCHAR,
		ChangDu:    ziFuShu,
		MoRenZhi:   FhDanYinHao + HFXXIAOXIE + FhDanYinHao,
	}
	return ret
}
