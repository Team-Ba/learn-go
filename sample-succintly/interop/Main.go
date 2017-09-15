package main

/*
  Answers the question, are all primitive int types interoperable? 
  The same test with primitive float types.
  Results in build error.
  
  Version: 20170910
*/
import (
  "fmt"
  "reflect"
)

var (
  f float32 = 7.0
  x uint = 7
  y int = 7
  z uint32 = 7
  a uint64 = 7
  b int32 = 7
  c int64 = 7
)

func main() {
  f2 := 7.0
  fmt.Printf("%f", f)
  fmt.Println(reflect.TypeOf(f))
  fmt.Printf("%f", f2)
  fmt.Println(reflect.TypeOf(f2))

  prodf := f * f2
  fmt.Printf("%f", prodf)
  fmt.Println(reflect.TypeOf(prodf))

  fmt.Printf("%d", x)
  fmt.Println(reflect.TypeOf(x))
  fmt.Printf("%d", y)
  fmt.Println(reflect.TypeOf(y))
  fmt.Printf("%d", z)
  fmt.Println(reflect.TypeOf(z))
  fmt.Printf("%d", a)
  fmt.Println(reflect.TypeOf(a))

  prod := x * y
  fmt.Printf("%d", prod)
  fmt.Println(reflect.TypeOf(prod))
}
