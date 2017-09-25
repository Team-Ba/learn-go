package main

import "fmt"

/*
  Tests struct methods with method receiver types
  
  Version: 20170925
*/
type rect struct {
  height int
  width int
}

// Creates copy of rect and operates on it
func (r rect) area() int {
  return r.height * r.width
}

// Gets rect pointer value and operates on it
func (r* rect) double() {
  r.height *= 2
  r.width *= 2
}

func main() {
  r1 := rect {
    height: 10,
    width: 20,
  }
  
  fmt.Printf("Height: %v \n", r1.height)
  fmt.Printf("Width: %v \n", r1.width)
  fmt.Printf("Area: %v \n", r1.area())
  
  fmt.Println()
  
  r1.double()
  fmt.Println("Doubled...")
  fmt.Printf("Height: %v \n", r1.height)
  fmt.Printf("Width: %v \n", r1.width)
  fmt.Printf("Area: %v \n", r1.area())
}