package main

import (
  "os"
  "fmt"
)

func main() {
  args := os.Args[1:]
  if len(args) == 0 {
    fmt.Println("Hello world!")
  } else if len(args) == 1 {
    fmt.Printf("Hello %s!\n", args[0])
  } else {
    fmt.Println("hello: too many parameters")
    os.Exit(1)
  }
}
