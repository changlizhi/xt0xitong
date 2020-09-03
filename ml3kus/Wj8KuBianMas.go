package ml3kus

//不能在初始化时进行因为很可能还没有这个表数据
// func init() {
// 	huanCunSuoYouJieGou() //初始化时用test作为初始化链接
// }
import (
	"log"
)

var (
	kuBianMa = map[string]interface{}{}
)

//这个方法唯一要注意的就是表和字段增到多大的时候会大到内存无法支持,当需要再次新增字段时只是在用户增加了字段的情况下需要再次查库更新
func huanCunBianMa() {
	suoShiLi.Do(func() { //这里需要把已存在的都纳入进来，所以需要新建一个配置文件，这个配置文件用go写成
		GengXinBianMa()
	})
}

//创建系统时会更新此数据
func GengXinBianMa() {
	log.Println("执行更新全局编码")
}
