package main

import "fmt"

func main(){
  var sum int = 100
  for i:= 1; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
