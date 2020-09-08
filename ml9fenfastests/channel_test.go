package ml9fenfastests

import(
  // "log"
  "testing"
)

func TestChannel1(t *testing.T){
  c := make(chan int)
  // go func(){
  //   x:= <- c
  //   y:= <- c
  //   // z:= <- c
  //   // g:= <- c
  //   log.Println("x---",x)
  //   log.Println("y---",y)
  //   // log.Println("z---",z)
  //   // log.Println("g---",g)
  // }()
    c <- 1
    // c <- 2
    // c <- 3
    // c <- 4
}

