package main

import "fmt"

/*
  Tests switch-case structure with fallthrough keyword.
  
  version: 20170916
*/
func main() {
  
  // Normal switch (Java perspective)
  switch y := "Go"; y {
      case "Python" :
        fmt.Println("Snake")
      
      case "Ruby" :
        fmt.Println("Gem")
      
      case "Go" :
        fmt.Println("Verb")
      
      case "C" :
        fmt.Println("Letter")
      
      case "Java" :
        fmt.Println("Coffee")
      
      default :
        fmt.Println("Word?")
  }
  
  // Abnormal switch (Java perspective)
  x := 5;
  switch {
    case x <= 5 :
      fmt.Println("x < 5")
      fallthrough
      
    case x >= 5 :
      fmt.Println("x == 5")
      fallthrough
    
    case x < 5 :
      fmt.Println("x < 5")
      
    case x == 5 :
      fmt.Println("x == 5")
      
    case x > 5 :
      fmt.Println("x == 5")
    
    default :
      fmt.Println("How'd that manage?");
  }
}