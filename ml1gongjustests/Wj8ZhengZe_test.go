package ml1gongjustests

import (
	"log"
	"testing"
	"xt0xitong/ml0gongjus"
)

func TestShouJiHaoPiPei(t *testing.T) {
	ret := ml0gongjus.ShouJiHaoPiPei("1112222ff3333")
	log.Println(ret, len(ret))
}
