package main

import (
  "strings"
  "code.google.com/p/go-tour/wc"
)



func WordCount(s string) map[string]int {
  m := make(map[string]int)
  var a []string = strings.Split(s, " ")
  for i:=0; i < len(a);i++{
    if m[a[i]] == 0 {
      m[a[i]] = 1
    }else{
      m[a[i]] += 1
    }
  }
  return m
}

func main() {
  wc.Test(WordCount)
}
