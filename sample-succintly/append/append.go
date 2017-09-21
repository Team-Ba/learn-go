package main

/*
  Tests slice and what will happen if:
  1. slice is initialized with value and then, appended
  2. slice is appended even if maximum capacity is reached
  
  Version: 20170921
*/
import "fmt" 

func main() {
  // Initialize slice with maximum of 8 ints and four 0's
  mySlice := make([]int, 4, 8)
  fmt.Println(mySlice)
  
  // Add values to slice
  mySlice = append(mySlice, 1, 2, 3, 4)
  fmt.Println(mySlice)
  
  // Add more values to slice even if it has reached more than maximum capacity
  mySlice = append(mySlice, 5, 6, 7)
  fmt.Println(mySlice)
}