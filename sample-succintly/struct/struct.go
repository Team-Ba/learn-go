package main

import "fmt"

/*
  Tests normal struct creation
  
  Version: 20170925
*/
type rect struct {
  height int
  width int
}

func main() {

  // Struct creation 1
  var r1 *rect
  r1          = new(rect)
  r1.height   = 10
  r1.width    = 15
  
  fmt.Printf("First struct - height: %v, width: %v", r1.height, r1.width)
  fmt.Println()
  
  // Struct creation 2
  r2 := rect {
    height    : 12, 
    width     : 17,
  }
  
  fmt.Printf("Second struct - height: %v, width: %v", r2.height, r2.width)
  fmt.Println()
  
  // Struct creation 3
  r3 := &rect {12, 20} // height, width
  fmt.Printf("Third struct - height: %v, width %v", r3.height, r3.width)
  fmt.Println()
}