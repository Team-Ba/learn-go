package main

/*
  Tests pointer magic in go lang
  
  Version: 20170921
*/
import "fmt"

func main() {
  a, b := 20, 71
  
  // Variable a's value and memory address
  fmt.Println(a)
  fmt.Println(&a)
  
  // Just tried multiple variable assignment so...
  fmt.Println(&b)
  
  // Store variable a's pointer to variable p
  p := &a
  fmt.Println(p)
  fmt.Println(&p) // Print p's memory addresss
  
  *p = 21
  fmt.Println(&p) // Print p's memory addresss
  
  // Variable a's value changes... Like magic!
  fmt.Println(a)
}