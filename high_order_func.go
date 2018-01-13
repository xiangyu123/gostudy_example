package main

import "fmt"

func foo(x int) func (int) int  {
    p := func (y int) int {
      return x + y
    }
    return p
}

func main(){
  fmt.Println(foo(3)(5))
}
