package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/brkss/designpattrens.go/decorator"
	"github.com/brkss/designpattrens.go/generator"
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


  // generator 
  for i := range generator.Fib(100000){
    fmt.Println(i)
  }

}
