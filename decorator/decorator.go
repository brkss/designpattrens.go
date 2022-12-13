package decorator

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)


type piFunc func(n int)(float64)

func PiWrapLogger(fun piFunc, logger *log.Logger) piFunc {
  return func(n int) (result float64){
    fn := func(n int) (result float64){
      defer func(t time.Time){
        logger.Printf("took=%v, n=%v, result=%v\n", time.Since(t), n, result)
      }(time.Now())
      return fun(n)
    }
    return fn(n)
  }
}

func PiWrapCache(fun piFunc, cache *sync.Map) piFunc {
  return func(n int)(float64){
    fn := func(n int)(float64){
      key := fmt.Sprintf("n=%d", n);
      val, ok := cache.Load(key);
      if ok {
        return val.(float64);
      }
      result := fun(n);
      cache.Store(key, result)
      return result
    }
    return fn(n)
  }
}

func Pi(n int) float64 {
  
  ch := make(chan float64)

  for k := 0; k <= n; k++ {
    go func(ch chan float64, k float64){
      ch <- 4 * math.Pow(-1, k) / (2 * k + 1)
    }(ch, float64(k))
  }

  result := 0.0
  for k := 0; k <= n; k++ {
    result += <- ch
  }

  return result
}


