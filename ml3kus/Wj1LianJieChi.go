package ml3kus

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"xt0xitong/ml2changliangs"
)

func init() {
	chuShiHuaChi() //初始化时用test作为初始化链接
}

var (
	suoShiLi sync.Once
	chi      = map[string]*sql.DB{}
)

func chuShiHuaChi() {
	suoShiLi.Do(func() { //这里需要把已存在的都纳入进来，所以需要新建一个配置文件，这个配置文件用go写成
		testDb := chuangJianChi(ml2changliangs.TEST)
		chi[ml2changliangs.TEST] = testDb
		// 除了XT0XITONG之外所有的业务系统都不需要在启动时创建数据库，因为有系统表，在用户手动添加的时候会自行创建库和表,添加之后就是自动添加数据库的基础表。
		ChuangJianJiChuBiao()
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
	if err == nil {
		chi[shuJuKuMing] = nil
		log.Println("删除数据库:", sqlStr, result, err, chi)
	}
}

func TianJiaLianJieChi(shuJuKuMing string) {
	if chi[shuJuKuMing] != nil {
		log.Printf("已存在，请直接使用:chi[%s]=", shuJuKuMing, chi[shuJuKuMing])
		return
	}
	ChuangJianKu(shuJuKuMing)
	lianJieChi := chuangJianChi(shuJuKuMing)
	chi[shuJuKuMing] = lianJieChi
}

//这里就是基础系统的数据结构，初始化系统时必须添加。
//字段里不需要挂上系统和表，因为除非把字段作为第一层，否则就没有意义，如果要维护两个对应关系就丧失了配置json的简便性要求，所以还是等做缓存的时候再把字段做成一个大对象map，或者还是直接循环出来就好，本来也就几k而已。基础表不用太多资源支持。现在就是写出一个基本逻辑，生成数据的时候要用表名加上字段的方式入值
//重新思考一下这个设计，针对的就是一个个的字段，根本不用在系统的基础上做这种限制，因为所有的系统都只有两种表，主键表和值表。
//初始化查询时应该以操作库为查询条件查到所有的表，然后再查出所有的字段，所以基础表是省不了的，因为没有初始条件了就没法拿到基础结构了。用太多if就丧失了使用配置的意义。

func ChuangJianJiChuBiao() {
	//读取json文件，拿到XT0XITONG的数据
	f, err := os.Open("../jichu.json")
	if err != nil {
		log.Println("jichu.json路径错误", err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("jichu.json读取错误", err)
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data := map[string]interface{}{}
	err = json.Unmarshal(fd, &data)
	if err != nil {
		log.Println("jichu.json解析错误", err)
	}

	for k, v := range data {
		if k == ml2changliangs.XT0XITONG { //其他库等用户自行添加
			TianJiaLianJieChi(k)
			i := ml2changliangs.Sz0
			for biaok, biaov := range v.(map[string]interface{}) {
				vduanyan := biaov.(map[string]interface{})

				canShu := map[string]interface{}{
					ml2changliangs.CaoZuoKu:   k,
					ml2changliangs.CaoZuoBiao: biaok,
					ml2changliangs.ZhuJian:    vduanyan[ml2changliangs.ZhuJian],
					ml2changliangs.SuoYin:     vduanyan[ml2changliangs.SuoYin],
					ml2changliangs.ZiDuans:    vduanyan[ml2changliangs.ZiDuans],
				}
				log.Println("i---", i, vduanyan[ml2changliangs.ZiDuans].([]interface{}))
				ChuangJianBiao(canShu)
				i++
			}
		}

	}
}
