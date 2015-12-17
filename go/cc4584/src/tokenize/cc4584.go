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
  left []string
  right []string
  LHS string
  RHS string
)

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

    typeTokenizer(LHS, RHS)

    for i:= 0; i < len(left); i++ {
      fmt.Println(left[i], "    ", right[i])
    }
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

  /* Primitive always at the right side
     VarType always at the left side
     FuncType couldn't match ListTyp */

  /* TypeVar can Unify with every Type */
  if isValidTypeVar(LHS) {
    left = append(left, LHS)
    right = append(right, RHS)
  } else if isValidTypeVar(RHS) {
    left = append(left, RHS)
    right = append(right, LHS)
  } else if isValidPrimitive(LHS) {

  /* Primitive can only Unify with itself */
    if LHS != RHS {
      fmt.Printf("%s\n", "1BOTTOM")
      os.Exit(1)
    }
  } else if isValidFuncType(LHS) {

  /* FuncType can only Unify with FuncType */
    if !isValidFuncType(RHS) {
      fmt.Printf("%s\n", "1BOTTOM")
      os.Exit(1)
    }

    /* Split it to Input Part and Output Part */
    var stack int = 0
    for i = 0 ; i < len(LHS); i++ {
      if LHS[i] == '(' {
        stack++
      }
      if LHS[i] == ')' {
        stack--
      }
      if stack == 0 {
        break
      }
    }

    tmp_left = strings.TrimSpace(LHS[0 : i+1])
    LHS = strings.TrimSpace(LHS[i+1 : len(LHS)])
    LHS = LHS[2: len(LHS)]

    for i = 0 ; i < len(RHS); i++ {
      if RHS[i] == '(' {
        stack++
      }
      if RHS[i] == ')' {
        stack--
      }
      if stack == 0 {
        break
      }
    }

    tmp_right = strings.TrimSpace(RHS[0 : i+1])
    RHS = strings.TrimSpace(RHS[i+1 : len(RHS)])
    RHS = RHS[2:len(RHS)]

    /* Unify output type */
    typeTokenizer(LHS, RHS)

    /* Remove "(" and ")" */
    tmp_left = strings.TrimSpace(tmp_left[1: len(tmp_left)-1])
    tmp_right = strings.TrimSpace(tmp_right[1:len(tmp_right)-1])

    /* ARGLIST is Empty */
    if len(tmp_left) == 0 && len(tmp_right) == 0 {
      fmt.Printf("%s\n", "2BOTTOM")
      os.Exit(1)
    }

    /* ARGLIST should not empty */
    if len(tmp_left) == 0 || len(tmp_right) == 0 {
      fmt.Printf("%s\n", "3BOTTOM")
      os.Exit(1)
    }

    /* Split ARGLIST Part */
    var bracket int = 0
    var start int = 0

    for i:=0 ; i < len(tmp_left); i++ {
      if tmp_left[i] == '(' {
        bracket ++
      }
      if tmp_left[i] == ')' {
        bracket--
      }
      if tmp_left[i] == ',' && bracket == 0 {
        tmp := strings.TrimSpace(tmp_left[start:i])
        start = i+1
        list_left = append(list_left, tmp)
      }
    }
    list_left = append(list_left, tmp_left[start: len(tmp_left)])

    start = 0
    for i:=0 ; i < len(tmp_right); i++ {
      if tmp_right[i] == '(' {
        bracket ++
      }
      if tmp_right[i] == ')' {
        bracket--
      }
      if tmp_right[i] == ',' && bracket == 0 {
        tmp := strings.TrimSpace(tmp_right[start:i])
        start = i+1
        list_right = append(list_right, tmp)
      }
    }
    list_right = append(list_right, tmp_right[start: len(tmp_right)])

    /* ARGLIST Should have same number of Type */
    if len(list_left) != len(list_right) {
      fmt.Printf("%s\n", "4BOTTOM")
      os.Exit(1)
    }

    /* Each part of ARGLIST should be Unified */
    for j := 0; j < len(list_left) ; j++ {
      typeTokenizer(list_left[j], list_right[j])
    }
  } else if !isValidList(RHS) {
  /* List can be only unified with List */
    fmt.Printf("%s\n", "5BOTTOM")
    os.Exit(1)
  } else {
    LHS = LHS[1: len(LHS)-1]
    RHS = RHS[1: len(RHS)-1]
    typeTokenizer(LHS, RHS)
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

  var stack int = 0
  var i int = 0
  for ; i < len(s); i++ {
    if s[i] == '(' {
      stack++
    }
    if s[i] == ')' {
      stack--
    }
    if stack < 0 {
      return false
    }
    if stack == 0 {
      break
    }
  }
  if i == len(s) {
    return false
  }

  /* Character after complete "()" should be "->" */
  right_part := strings.TrimSpace(s[i+1: len(s)])

  if right_part[0:2] != "->" {
    return false
  }

  /* String after "->" should be valid TYPE */
  if ! isValidType(right_part[2: len(right_part)]) {
    return false
  }

  s = strings.TrimSpace(s[0:i+1])

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

  /* Split by outside bracket",", every part should be valid TYPE */
  var bracket int = 0
  var start int = 0

  for i:=0 ; i < len(s); i++ {
    if s[i] == '(' {
      bracket ++
    }
    if s[i] == ')' {
      bracket--
    }
    if s[i] == ',' && bracket == 0 {
      tmp := strings.TrimSpace(s[start:i])
      start = i+1
      if !isValidType(tmp) {
        return false
      }
    }
  }
  if bracket != 0 {
    return false
  }
  if !isValidType(s[start: len(s)]) {
    return false
  }

  return true
}
