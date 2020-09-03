package ml4kustests

import (
	"testing"
	"xt0xitong/ml3kus"
	"xt0xitong/ml2changliangs"
)

//在这里把表和字段数据都添加了
func TestXinZeng(t *testing.T) {
	//数据库会初始化所以不需要自己操作，但是表不会，所以还要创建数据表之后再进行新增，新增完之后删除数据库
	ml3kus.HuoQuLianJieChi(ml2changliangs.XT0XITONG) //初始化的时候基础表就已经创建了
	//先把基础表的数据入库。重点需要入库的有bm0xitongs,bm1biaomings
  canShu := map[string]interface{}{
  	ml2changliangs.CaoZuoKu:   ml2changliangs.XT0XITONG,
  	ml2changliangs.CaoZuoBiao: ml2changliangs.Bm1Biaos, //Bm1Biaos中字段定义了多少字段就要声明多少字段。
    ml2changliangs.CaoZuoZhis: map[string]interface{}{
      ml2changliangs.BianMa:      "Ywb1YongHus",
      ml2changliangs.MingCheng:   "用户",
      ml2changliangs.ZhuJianBiao: ml2changliangs.ZjBiao + ml2changliangs.Zf1 + ml2changliangs.XiaoXies,
    },
  }
  ml3kus.XinZeng(canShu)
}
