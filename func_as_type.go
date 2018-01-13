package main

import "fmt"

type A func(int, int)

fmt.Printf("%T\n", A)

func (f A)Serve() {
    fmt.Println("serve2")
}

func serve(a int,b int) {
    c := a + b
    fmt.Println("serve1")
    fmt.Println(c)
}

func main() {
    a := A(serve)
    fmt.Printf("%T\n", a)
    a(1,2)
    a.Serve()
}
