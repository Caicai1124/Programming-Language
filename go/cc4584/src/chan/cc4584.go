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

var ch_left chan string = make(chan string)
var ch_right chan string = make(chan string)

func main(){
  inputReader = bufio.NewReader(os.Stdin)

  /*
  m := make(map[string]string)
  var elem string
  */
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

    go typeTokenizer(LHS, RHS)

    fmt.Printf("%s\n", <- ch_left)
    fmt.Printf("%s\n", <- ch_right)
    fmt.Printf("%s\n", <- ch_left)
    fmt.Printf("%s\n", <- ch_right)
    fmt.Printf("%s\n", <- ch_left)
    fmt.Printf("%s\n", <- ch_right)

   break
    /* Primitive always at the right side
       VarType always at the left side
       FuncType couldn't match ListTyp */

  }
}

func typeTokenizer(LHS string, RHS string){
  var (
    i int
    tmp_left string
    tmp_right string
    list_left []string
    list_right []string
  )

  LHS = strings.TrimSpace(LHS)
  RHS = strings.TrimSpace(RHS) 

  if isValidPrimitive(LHS) {
    ch_left <- RHS
    ch_right <- LHS
  } else if isValidPrimitive(RHS) {
    ch_left <- LHS
    ch_right <-RHS
  } else if isValidTypeVar(LHS) {
    ch_left <- LHS
    ch_right <- RHS
  } else if isValidTypeVar(RHS) {
    ch_left <- RHS
    ch_right <- LHS
  } else if isValidFuncType(LHS) {
    if !isValidFuncType(RHS) {
      fmt.Println("WRONG")
    }
    /* Unify output type */
    i = strings.LastIndex(LHS, "->")
    tmp_left = strings.TrimSpace(LHS[0 : i])
    LHS = strings.TrimSpace(LHS[i+2 : len(LHS)])

    i = strings.LastIndex(RHS, "->")
    tmp_right = strings.TrimSpace(RHS[0 : i])
    RHS = strings.TrimSpace(RHS[i+2 : len(RHS)])

    go typeTokenizer(LHS, RHS)

    /* Remove "(" and ")" */
    tmp_left = strings.TrimSpace(tmp_left[1: len(tmp_left)-1])
    tmp_right = strings.TrimSpace(tmp_right[1:len(tmp_right)-1])

    if len(tmp_left) == 0 && len(tmp_right) == 0 {
      fmt.Println("TEST1")
    }

    if len(tmp_left) == 0 || len(tmp_right) == 0 {
      fmt.Println("TEST2")
    }

    list_left = strings.Split(tmp_left, ",")
    list_right = strings.Split(tmp_right, ",")

    if len(list_left) != len(list_right) {
      fmt.Printf("%s\n", "BOTTOM")
      os.Exit(1)
    }

    for j := 0; j < len(list_left) ; j++ {
      go typeTokenizer(list_left[j], list_right[j])
    }
  } else if !isValidList(RHS) {
    fmt.Println("TEST3")
  } else {
    LHS = LHS[1: len(LHS)-1]
    RHS = RHS[1: len(RHS)-1]
    go typeTokenizer(LHS, RHS)
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
  /* begin with "`" */
  if !strings.HasPrefix(s, "`") {
    return false
  }

  /* At least one letter */
  if len(s) < 2 {
    return false
  }

  /* First character should be letter */
  if !isLetter(s[1:2]) {
    return false
  }

  /* Rest Character should be letter or digit */
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
  if !strings.HasPrefix(s, "(") {
    return false
  }
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
  if !strings.HasPrefix(s, "[") {
    return false
  }
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
