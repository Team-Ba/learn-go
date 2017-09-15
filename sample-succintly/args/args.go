package main

/*
  Test args in Go lang.
  os.Args[0] is name of path/build running this.
  If run in Windows7 using go run, %USERS%\AppData\Local\Temp\go-build<bunch of numbers>\command-line-arguments\_obj\exe\args.exe <arguments>
  If run using exe build, <path where args is executed>\args <arguments>
  
  Version: 20170916
*/
import (
  "fmt"
  "os"
)

func main() {
  lenArgs := len(os.Args)
  
  // For
  for i := 1; i < lenArgs; i++ {
    fmt.Println(os.Args[i])
  }
  
  // For each
  for _, arg := range os.Args {
    fmt.Println(arg)
  }
  fmt.Println(os.Args[0])
}