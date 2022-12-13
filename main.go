package main

import (
	"log"
	"os"
	"sync"

	"github.com/brkss/designpattrens.go/decorator"
)


func main(){

  // decorator !
  // l( c( pi() ) )
  c := decorator.PiWrapCache(decorator.Pi, &sync.Map{});
  l := decorator.PiWrapLogger(c, log.New(os.Stdout, "test ", 1));

  
  l(10000);
  l(400000);
  l(10000);
  l(10000);

  //fmt.Println(decorator.Pi(1000))
  //fmt.Println(decorator.Pi(50000))

}
