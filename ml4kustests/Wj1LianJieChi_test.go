package ml4kustests

import (
	"database/sql"
	"log"
	"testing"
	"time"
	"xt0xitong/ml2changliangs"
	"xt0xitong/ml3kus"
)

func TestHuoQuLianJieChi(t *testing.T) {
	count := make(chan int, 1)
	db := ml3kus.HuoQuLianJieChi(ml2changliangs.XT1YONGHU)[ml2changliangs.Ceng1].(*sql.DB)
	log.Println("TestHuoQuLianJieChi,db:", db)
	for i := 0; i < 10000; i++ { //十万个并发去拿相同的连接池都不会报错，下一步是真正建立链接进行查询，可以校验maxConn是否设置成功
		go func(i int) {
			count <- i
			err := db.Ping()
			log.Println("db---", err, i, db)
			<-count
		}(i)
	}
	time.Sleep(time.Duration(5) * time.Hour)
}
