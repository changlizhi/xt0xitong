package ml3kus

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
	"sync"
	"xt0xitong/ml2changliangs"
)

func init() {
	testDb := chuangJianChi(ml2changliangs.TEST)
	chi[ml2changliangs.TEST] = testDb
	chuShiHuaChi() //初始化时用test作为初始化链接
}

var (
	suoShiLi sync.Once
	chi      = map[string]*sql.DB{}
)

func chuShiHuaChi() {
	suoShiLi.Do(func() { //这里需要把已存在的都纳入进来，所以需要新建一个配置文件，这个配置文件用go写成
		for _, v := range strings.Split(ml2changliangs.Kus, ml2changliangs.FhDouHao) {
			TianJiaLianJieChi(v)
		}
	})
}

func chuangJianChi(shuJuKuMing string) *sql.DB {
	lianJieChi, err := sql.Open("mysql", "root:rootclz@tcp(127.0.0.1:3306)/"+shuJuKuMing) //进入mysql数据库，然后
	if err != nil {
		log.Println("建立链接失败", err)
		os.Exit(1)
	}
	lianJieChi.SetMaxOpenConns(50)
	lianJieChi.SetMaxIdleConns(5)
	chi[shuJuKuMing] = lianJieChi
	log.Println("chuangJianChi:lianJieChi---", lianJieChi)
	return lianJieChi
}

func HuoQuLianJieChi(shuJuKuMing string) map[string]interface{} {
	ret := map[string]interface{}{
		ml2changliangs.Ceng1: chi[shuJuKuMing],
	}
	return ret
}

func ChuangJianKu(shuJuKuMing string) {
	db := chi[ml2changliangs.XiTongKu] //用test库获取链接
	sqlStr := "CREATE DATABASE " + shuJuKuMing
	result, err := db.Exec(sqlStr) //这里是为了避免无数据库的情况发生，做一个容错
	log.Println("创建新的数据库:", sqlStr, result, err, chi)
}
func ShanChuKu(shuJuKuMing string) {
	db := chi[ml2changliangs.XiTongKu] //用test库获取链接
	sqlStr := "DROP DATABASE " + shuJuKuMing
	result, err := db.Exec(sqlStr) //这里是为了避免无数据库的情况发生，做一个容错
  if err == nil{
    chi[shuJuKuMing] = nil
    log.Println("删除数据库:", sqlStr, result, err, chi)
  }
}

func ShanChuJiChuKu() {
	for _, v := range strings.Split(ml2changliangs.Kus, ml2changliangs.FhDouHao) {
		ShanChuKu(v)
	}
}

func TianJiaLianJieChi(shuJuKuMing string) {
	if chi[shuJuKuMing] != nil {
		log.Printf("已存在，请直接使用:chi[%s]=", shuJuKuMing,chi[shuJuKuMing])
    return
	}
	ChuangJianKu(shuJuKuMing)
	lianJieChi := chuangJianChi(shuJuKuMing)
	chi[shuJuKuMing] = lianJieChi
}
