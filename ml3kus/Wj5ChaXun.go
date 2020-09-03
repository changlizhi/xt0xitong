package ml3kus

import (
	"database/sql"
	"log"
	"sort"
	"strings"
	"xt0xitong/ml2changliangs"
)

func ChaXun(canShu map[string]interface{}) map[string]interface{} {
	//select ziduans from caoZuoKu.caoZuoBiao where ziDuan=?
	caoZuoKu := canShu[ml2changliangs.CaoZuoKu].(string)
	caoZuoBiao := canShu[ml2changliangs.CaoZuoBiao].(string)
	ziDuans := canShu[ml2changliangs.ZiDuans].([]interface{})                        //把要查询的字段拿出来
	tiaoJianHeZhis := canShu[ml2changliangs.TiaoJianHeZhis].(map[string]interface{}) //把字段拿出来

	chaXunZiDuans := []string{}
	for _, v := range ziDuans {
		chaXunZiDuans = append(chaXunZiDuans, v.(string))
	}
	sort.Strings(chaXunZiDuans)
	wenHaos := []string{}
	values := []interface{}{}
	for k, v := range tiaoJianHeZhis {
		wenHaos = append(wenHaos, k+ml2changliangs.FhDengHao+ml2changliangs.FhWenHao)
		values = append(values, v)
	}

	builder := strings.Builder{}
	builder.WriteString(" SELECT ")
	builder.WriteString(strings.Join(chaXunZiDuans, ml2changliangs.FhDouHao))
	builder.WriteString(" FROM ")
	builder.WriteString(caoZuoKu)
	builder.WriteString(ml2changliangs.FhDianHao)
	builder.WriteString(caoZuoBiao)
	if len(wenHaos) > ml2changliangs.Sz0 {
		builder.WriteString(" WHERE ")
		builder.WriteString(strings.Join(wenHaos, " AND "))
	}

	sqlStr := builder.String()
	db := HuoQuLianJieChi(caoZuoKu)[ml2changliangs.Ceng1].(*sql.DB)

	rows, err := db.Query(sqlStr, values...)
	//在没有报错的情况下可以执行Next和Scan方法把所有查到的字段值拿出来放到CanShu里返回回去
	retShuJu := scanRet(chaXunZiDuans, rows)
	ret := map[string]interface{}{
		ml2changliangs.Ceng1: retShuJu,
	}
	log.Println("ChaXun:sqlStr,values,err,ret---", sqlStr, values, err, ret)
	return ret
}

func scanRet(lies []string, rows *sql.Rows) []map[string]interface{} {
	//必须保证查到的列顺序和传入的列顺序一致，否则将会匹配失败
	dbCols, err := rows.Columns()
	if err != nil {
		log.Println("查询数据库失败，请注意检查！")
	}
	for i, _ := range dbCols {
		if dbCols[i] != lies[i] {
			log.Println("数据库返回的字段顺序和查询时字段顺序不一致，请注意检查")
		}
	}
	//保证了顺序之后再Scan就没问题了
	ret := []map[string]interface{}{}
	for rows.Next() { //每一行
		tempLie := make([]string, len(lies))
		tempLieRef := make([]interface{}, len(lies))

		for i := ml2changliangs.Sz0; i < len(lies); i++ {
			tempLieRef[i] = &tempLie[i]
		}
		rows.Scan(
			tempLieRef...,
		)
		retOne := map[string]interface{}{}
		for xiaBiao, lie := range lies {
			if tempLieRef[xiaBiao] != nil {
				retOne[lie] = *tempLieRef[xiaBiao].(*string)
			}
		}
		ret = append(ret, retOne)
	}
	return ret
}
