package ml1gongjustests

import (
	"log"
	"testing"
	"xt0xitong/ml0gongjus"
)

func TestDuQuMoXingJson(t *testing.T) {
	ret := ml0gongjus.DuQuMoXingJson()
	for _, r := range ret {
		log.Println("Wj10MoxingJson_test.go---TestDuQuMoXingJson,r", r)
	}
}
