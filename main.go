package main

import (
	"log"
	"os"

	"github.com/brkss/designpattrens.go/decorator"
)


func main(){

  // decorator !
  f := decorator.PiWrapLogger(decorator.Pi, log.New(os.Stdout, "test ", 1));

  f(10000);

  //fmt.Println(decorator.Pi(1000))
  //fmt.Println(decorator.Pi(50000))

}
