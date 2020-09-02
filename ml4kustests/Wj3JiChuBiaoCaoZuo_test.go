package ml4kustests

import (
	"testing"
	"xt0xitong/ml3kus"
	"xt0xitong/ml2changliangs"
)

func TestChuangJianJiChuBiaos(t *testing.T) {
	ml3kus.HuoQuLianJieChi(ml2changliangs.XT0XITONG)//初始化的时候基础表就已经创建了
	ml3kus.ShanChuKu(ml2changliangs.XT0XITONG) //测试完成后删除数据库，基础测试在所有的测试方法中都不需要前置测试存在。
}
