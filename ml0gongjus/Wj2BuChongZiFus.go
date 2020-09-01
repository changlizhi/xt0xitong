package ml0gongjus

import (
	"strings"
	"xt0xitong/ml2changliangs"
)

func BuChongChuan(geShu int, ziFu string) string {
	if geShu < ml2changliangs.Sz1 {
		return ml2changliangs.FhKongZiFu
	}
	retshuzu := []string{}
	for i := ml2changliangs.Sz0; i < geShu; i++ {
		retshuzu = append(retshuzu, ziFu)
	}
	return strings.Join(retshuzu, ml2changliangs.FhKongZiFu)
}

func ZuoBuChong(geShu int, ziFu string, daiBuChong string) string {
	buChong := BuChongChuan(geShu, ziFu)
	ret := buChong + daiBuChong
	return ret
}
func YouBuChong(geShu int, ziFu string, daiBuChong string) string {
	buChong := BuChongChuan(geShu, ziFu)
	ret := daiBuChong + buChong
	return ret
}
func ZuoBuQi(zongWeiShu int, ziFu string, daiBuChong string) string {
	daiBuChongWeiShu := len(daiBuChong)
	if daiBuChongWeiShu >= zongWeiShu {
		return daiBuChong
	}
	geShu := zongWeiShu - daiBuChongWeiShu
	buChong := BuChongChuan(geShu, ziFu)
	ret := buChong + daiBuChong
	return ret
}
func YouBuQi(zongWeiShu int, ziFu string, daiBuChong string) string {
	daiBuChongWeiShu := len(daiBuChong)
	if daiBuChongWeiShu >= zongWeiShu {
		return daiBuChong
	}
	geShu := zongWeiShu - daiBuChongWeiShu
	buChong := BuChongChuan(geShu, ziFu)
	ret := daiBuChong + buChong
	return ret
}

func ShuZiZiFuBu1Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz1, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu2Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz2, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu3Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz3, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu4Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz4, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu5Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz5, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu6Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz6, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu7Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz7, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu8Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz8, ml2changliangs.Zf0, ziFu)
}
func ShuZiZiFuBu9Ge0(ziFu string) string {
	return YouBuChong(ml2changliangs.Sz9, ml2changliangs.Zf0, ziFu)
}
