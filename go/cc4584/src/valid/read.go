package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

var (
  inputReader *bufio.Reader
  input string
  err error 
  strList []string
  LHS string
  RHS string
)

func main(){
  inputReader = bufio.NewReader(os.Stdin)

  for {
    /* Get input from Screen */
    input, err = inputReader.ReadString('\n')

    /* Strip last "\n" */
    input = strings.TrimSpace(input)

    /* Split at "&" */
    strList = strings.Split(input, "&")

    /* Deal with "QUIT" */
    if len(strList) == 1 {
      if strings.Trim(strList[0]," ") == "QUIT" {
        os.Exit(1)
      } else {
        fmt.Printf("%s\n", "ERR")
        os.Exit(1)
      }
    } else if len(strList) != 2 {
      fmt.Printf("%s\n", "ERR")
      os.Exit(1)
    }

    LHS = strings.TrimSpace(strList[0])
    RHS = strings.TrimSpace(strList[1])

    if isValidType(LHS) != true || isValidType(RHS) != true{
      fmt.Printf("%s\n", "ERR")
      os.Exit(1)
    }
  }
}

/* Four kinds of Types */
func isValidType(s string) bool {
  s = strings.TrimSpace(s)
  if len(s) == 0 {
    return false 
  }

  switch s[0:1] {
    case "`" : return isValidTypeVar(s)
    case "(" : return isValidFuncType(s)
    case "[" : return isValidList(s)
    default : return isValidPrimitive(s)
  }
}

func isValidTypeVar(s string) bool{
  if len(s) < 2 {
    return false
  }

  if !isLetter(s[1:2]) {
    return false
  }
  for i := 2; i < len(s); i++ {
    if !isLetter(s[i:i+1]) && !isDigit(s[i:i+1]) {
      return false
    }
  }
  return true
}

func isLetter(s string) bool{
  if strings.Compare(s,"A") >= 0 && strings.Compare(s, "Z") <=0 {
    return true
  }

  if strings.Compare(s,"a") >= 0 && strings.Compare(s, "z") <=0 {
    return true
  }
  return false
}

func isDigit(s string) bool{
  if strings.Compare(s,"0") >= 0 && strings.Compare(s, "9") <=0 {
    return true
  }
  return false
}

func isValidFuncType(s string) bool{
  last := strings.LastIndex(s, "->")

  /* String after "->" should be valid TYPE */
  if ! isValidType(s[last+2:len(s)]) {
    return false
  }

  s = strings.TrimSpace(s[0:last])

  /* Should start with "(" and end with ")" */
  if !strings.HasSuffix(s, ")") {
    return false
  }

  return isValidArgList(s[1:len(s)-1])
}

func isValidList(s string) bool{
  if !strings.HasSuffix(s, "]") {
    return false
  }
  return isValidType(s[1:len(s)-1])
}

func isValidPrimitive(s string) bool{
  if s == "int" || s == "real" || s == "str"{
    return true
  }

  return false
}

func isValidArgList(s string) bool{
  s = strings.TrimSpace(s)
  if len(s) == 0 {
    return true
  }

  /* Split by ",", every part should be valid TYPE */
  list := strings.Split(s, ",")

  for i:= 0; i < len(list);i++ {
    list[i] = strings.TrimSpace(list[i])
    if len(list[i]) == 0 {
      return false
    }
    if !isValidType(list[i]) {
      return false
    }
  }
  return true
}
