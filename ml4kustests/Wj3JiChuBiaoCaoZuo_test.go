package ml4kustests

import (
	"testing"
	"xt0xitong/ml2changliangs"
	"xt0xitong/ml3kus"
)

func TestChuangJianJiChuBiaos(t *testing.T) {
	ml3kus.HuoQuLianJieChi(ml2changliangs.XT0XITONG) //初始化的时候基础表就已经创建了
	// ml3kus.ShanChuKu(ml2changliangs.XT0XITONG) //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
//从操作上来说创建了基础表就是为了添加业务表结构数据的，所以下一步就是根据需要从页面上添加表结构数据