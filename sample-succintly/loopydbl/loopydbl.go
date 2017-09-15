package main

/*
  Testing 2D array.
  Notes:
    range can only be used as part of expression (loop only, for now)
    range 1st value = index
    range 2nd value = element
  
  version: 20170916
*/
import "fmt"

func main() {
  var arr2 = [3][3]int {
    {1, 1, 1},
    {1, 2, 2},
    {1, 2, 3},
  }
  
  for i := 0; i < len(arr2); i++ {
    for x, y := range arr2[i] {
      fmt.Printf("x: %d\n", x)
      fmt.Println(y)
    }
  
    fmt.Println(arr2[i])
  }
}