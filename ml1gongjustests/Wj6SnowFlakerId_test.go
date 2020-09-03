package ml1gongjustests

import (
	"log"
	"testing"
	"time"
	"xt0xitong/ml0gongjus"
)

func TestHuoQuId(t *testing.T) {
	ret := ml0gongjus.HuoQuId()
	log.Println(ret)
}
func TestHuoQuIdZiFu(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go func(i int) {
			ret := ml0gongjus.HuoQuIdZiFu()
			log.Println("TestHuoQuIdZiFu---", i, ret, len(ret))
		}(i)
	}
	time.Sleep(time.Duration(5) * time.Hour)

}
