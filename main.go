package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/brkss/designpattrens.go/decorator"
	"github.com/brkss/designpattrens.go/generator"
	"github.com/brkss/designpattrens.go/observer"
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
 
  sub := observer.EventSubject{
    Observers: sync.Map{},
  }

  var obs1 = observer.EventObserver{Id: 1, Time: time.Now()}
  var obs2 = observer.EventObserver{Id: 2, Time: time.Now()}

  sub.AddListener(&obs1)
  sub.AddListener(&obs2)
  
  // remove observer after x time 
  go func(){
    select {
      case <- time.After(time.Microsecond * 10): {
        sub.RemoveListener(&obs1) 
      }
    }
  }()

  for i := range generator.Fib(100000){
    sub.Notify(observer.Event{Data: i})
    time.Sleep(time.Microsecond * 10)
    //fmt.Println(i)
  }

}
