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
//添加完结构数据之后就是生成结构表对应的各个主键表和字段值表，都是要在业务中体现的，每个系统各自有各自的业务表，基础表虽然结构完全相同但还是为了简单起见不拆分
//然后就是业务表数据的缓存更新和基本增删改查，每次新增字段之后缓存都要改
//业务数据真实业务的多条件展示和界面配合
//现在没有数据所以新增开始


