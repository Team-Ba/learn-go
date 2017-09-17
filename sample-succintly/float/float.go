package main

/*
  Tests what kind of float/int data type is automatically assigned when not defined

  Version: 20170910
*/
import (
  "fmt"
  "reflect"
)

var (
  x float32 = 7.0
)

func main() {
  y := 7.0
  z := 7
  
  fmt.Printf("%f", x)
  fmt.Println(reflect.TypeOf(x))
  
  fmt.Printf("%f", y)
  fmt.Println(reflect.TypeOf(y))
  
  fmt.Printf("%f", z)
  fmt.Println(reflect.TypeOf(z))
}
