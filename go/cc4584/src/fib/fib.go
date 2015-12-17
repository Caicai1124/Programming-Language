package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func(int) int {
  return func(x int) int{
    if x == 0 || x == 1 {
      return 1
    } else{
      t := fibonacci()
      return t(x-1)+t(x-2)
    }
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f(i))
  }
}
