package main

import "fmt"

func count() {
  for i:= 0; i < 10; i++ {
    defer fmt.Println(i)
  }
}
func main(){
  fmt.Println("counting")
  count()
  fmt.Println("DONE")
}
