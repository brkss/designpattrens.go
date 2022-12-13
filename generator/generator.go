
package generator



func Fib(n int) chan int {
  
  out := make(chan int)

  go func() {
    defer close(out) // close chanel after calculating fib to avoid deadlock
    for i, j := 0, 1; i < n; i, j = i+j, i {
      out <- i
    } 
  }()

  return out

}
