package ml3kus

import (
	"database/sql"
	"log"
	"strings"
	"xt0xitong/ml0gongjus"
	"xt0xitong/ml2changliangs"
)

//为了创建表时不报错，所以提供一个删除表和数据库的方法，测试创建表或者创建数据库的时候

//后续需要实现一下，在创建表时不做那么多限制

func SheZhiWeiYiSuoYin(canShu map[string]interface{}) map[string]interface{} {
	//ALTER TABLE XT0JICHU.bm1biaomings ADD UNIQUE INDEX BianMa (BianMa);
	caoZuoKu := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.CaoZuoKu])
	caoZuoBiao := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.CaoZuoBiao])
	suoYin := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.SuoYin])
	builder := strings.Builder{}

	builder.WriteString(" ALTER TABLE ")
	builder.WriteString(caoZuoKu)
	builder.WriteString(ml2changliangs.FhDianHao)
	builder.WriteString(caoZuoBiao)
	builder.WriteString(" ADD UNIQUE INDEX ")
	builder.WriteString(suoYin)
	builder.WriteString(" ( ")
	builder.WriteString(suoYin)
	builder.WriteString(" ) ")

	sqlStr := builder.String()
	db := HuoQuLianJieChi(caoZuoKu)[ml2changliangs.Ceng1].(*sql.DB)
	result, err := db.Exec(sqlStr)
	log.Println("SheZhiWeiYiSuoYin:sqlStr,result,err---", sqlStr, result, err)
	return canShu
}

func ChuangJianBiao(canShu map[string]interface{}) map[string]interface{} {
	//CREATE TABLE caoZuoKu.BiaoMing (
	// `ZhuJian` BIGINT(20) NOT NULL,
	// `MingCheng` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// `BianMa` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// `ZhuJianZhiBiao` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// PRIMARY KEY (`ZhuJian`)
	//  )
	//  ENGINE=InnoDB
	// 这里要求传入的ShuJu第一个必须是表名字符串，
	// 第二个必须是主键名指定
	// 第三个必须是字段列表
	// 后续强化这个校验
	// 要重点命名清楚每个中间变量的意义，ShuJuKu是需要访问的数据库，BiaoMing是需要访问的数据库下的表，字段中的ShuJuKu是值

	caoZuoKu := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.CaoZuoKu])
	caoZuoBiao := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.CaoZuoBiao])
	zhuJian := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.ZhuJian])
	suoYin := ml0gongjus.DuanYanZiFuChuan(canShu[ml2changliangs.SuoYin])
	ziDuans := canShu[ml2changliangs.ZiDuans].([]map[string]interface{}) //把字段拿出来

	builder := strings.Builder{}

	builder.WriteString("CREATE TABLE ")
	builder.WriteString(caoZuoKu)
	builder.WriteString(ml2changliangs.FhDianHao)
	builder.WriteString(caoZuoBiao)
	builder.WriteString(" (")
	for _, v := range ziDuans {
		//`MingCheng` VARCHAR(50) NOT NULL DEFAULT 'hfx',
		builder.WriteString(ml0gongjus.DuanYanZiFuChuan(v[ml2changliangs.ZiDuanMing]))
		builder.WriteString(" ")
		builder.WriteString(ml0gongjus.DuanYanZiFuChuan(v[ml2changliangs.LeiXing]))
		builder.WriteString("(")
		builder.WriteString(ml0gongjus.DuanYanZiFuChuan(v[ml2changliangs.ChangDu]))
		builder.WriteString(") NOT NULL DEFAULT ")
		builder.WriteString(ml0gongjus.DuanYanZiFuChuan(v[ml2changliangs.MoRenZhi]))
		builder.WriteString(",")
	}

	builder.WriteString("PRIMARY KEY (")
	builder.WriteString(zhuJian)
	builder.WriteString(")")
	if suoYin != ml2changliangs.FhKongZiFu {
		builder.WriteString(",UNIQUE INDEX ")
		builder.WriteString(suoYin)
		builder.WriteString(" (")
		builder.WriteString(suoYin)
		builder.WriteString(")")
	}
	builder.WriteString(")COLLATE='utf8mb4_general_ci' ENGINE=InnoDB")

	sqlStr := builder.String()

	db := HuoQuLianJieChi(caoZuoKu)[ml2changliangs.Ceng1].(*sql.DB)
	result, err := db.Exec(sqlStr)
	if err != nil && (strings.Contains(err.Error(), "1049") || strings.Contains(err.Error(), "1146")) {
		log.Println("数据库不存在", err)
		TianJiaLianJieChi(caoZuoKu)
		ChuangJianBiao(canShu)
	}
	log.Println("ChuangJianBiao:sqlStr,result,err---", sqlStr, result, err)
	return canShu
}
