package ml3kus

import (
	"database/sql"
	"log"
	"strings"
	"xt0xitong/ml0gongjus"
	"xt0xitong/ml2changliangs"
)

func XinZeng(canShu map[string]interface{}) map[string]interface{} {
	//insert into caoZuoKu.caoZuoBiao (columns) values(values)
	caoZuoKu := canShu[ml2changliangs.CaoZuoKu].(string)
	caoZuoBiao := canShu[ml2changliangs.CaoZuoBiao].(string)
	caoZuoZhis := canShu[ml2changliangs.CaoZuoZhis].(map[string]interface{}) //把字段拿出来

	keys := []string{}
	wenHaos := []string{}
	values := []interface{}{}
	//默认ZhuJian全都通过snowflakerid得到，但是业务表中主键是从主键表来的。
	if caoZuoZhis[ml2changliangs.ZhuJian] == nil {
		keys = append(keys, ml2changliangs.ZhuJian)
		values = append(values, ml0gongjus.HuoQuId())
		wenHaos = append(wenHaos, ml2changliangs.FhWenHao)
	}
	for k, v := range caoZuoZhis {
		keys = append(keys, k)
		values = append(values, v)
		wenHaos = append(wenHaos, ml2changliangs.FhWenHao)
	}

	builder := strings.Builder{}
	builder.WriteString(" INSERT INTO ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(ml2changliangs.FhDianHao)
	builder.WriteString(caoZuoBiao)

	builder.WriteString(" ( ")
	builder.WriteString(strings.Join(keys, ","))
	builder.WriteString(" )values( ")
	builder.WriteString(strings.Join(wenHaos, ","))
	builder.WriteString(" ) ")

	sqlStr := builder.String()
	db := HuoQuLianJieChi(caoZuoKu)[ml2changliangs.Ceng1].(*sql.DB)
	result, err := db.Exec(sqlStr, values...)
	log.Println("XinZeng:sqlStr,values,result,err---", sqlStr, values, result, err)
	return canShu
}
