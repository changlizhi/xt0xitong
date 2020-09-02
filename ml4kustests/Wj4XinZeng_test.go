package ml4kustests

import (
	"testing"
	"xt0xitong/ml3kus"
)

//在这里把表和字段数据都添加了
func TestXinZeng(t *testing.T) {
	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	ml3kus.ChuangJianJiChuBiao() //基础表必须存在，不需要单独创建
	//先把基础表的数据入库。重点需要入库的有bm0xitongs,bm1biaomings
}
