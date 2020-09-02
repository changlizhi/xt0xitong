package ml3kus

import (
	// "xt0xitong/ml2changliangs"
	"io/ioutil"
	"log"
	"os"
)

//这里就是基础系统的数据结构，初始化系统时必须添加。
//字段里不需要挂上系统和表，因为除非把字段作为第一层，否则就没有意义，如果要维护两个对应关系就丧失了配置json的简便性要求，所以还是等做缓存的时候再把字段做成一个大对象map，或者还是直接循环出来就好，本来也就几k而已。基础表不用太多资源支持。现在就是写出一个基本逻辑，生成数据的时候要用表名加上字段的方式入值
//重新思考一下这个设计，针对的就是一个个的字段，根本不用在系统的基础上做这种限制，因为所有的系统都只有两种表，主键表和值表。
//初始化查询时应该以操作库为查询条件查到所有的表，然后再查出所有的字段，所以基础表是省不了的，因为没有初始条件了就没法拿到基础结构了。用太多if就丧失了使用配置的意义。

func ChuangJianJiChuBiao() {
	//读取json文件，拿到XT0XITONG的数据
	f, err := os.Open("../jichu.json")
	if err != nil {
		log.Println("read file fail", err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
	}
	log.Println("fd---", string(fd))
}
