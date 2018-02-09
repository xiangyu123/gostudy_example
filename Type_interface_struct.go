package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Color string
}

type Meower interface {
	Meow(d string)
}

func Greet(meower Meower, d string) {
	meower.Meow(d)
}

func (c Cat) Meow(d string) {
	c.Name = d
}

func main() {
	a := Cat{"a", "black"}
	fmt.Println("Before Name: ", a.Name, "Color: ", a.Color)
	Greet(a, "aa")
	// a.Meow("aa")
	fmt.Println("After Name: ", a.Name, "Color: ", a.Color)
	b := &Cat{"b", "black"}
	fmt.Println("Before Name: ", b.Name, "Color: ", b.Color)
	Greet(b, "bb")
	fmt.Println("After Name: ", b.Name, "Color: ", b.Color)
}
// func (c Cat) Neow(d string) 实现了Meower接口，由于接受者Cat是值类型，因此会隐式生成一个*Cat的同名函数，因此*Cat 也实现了Meower接口，但是反过来
//接受者是*Cat类型却不会生成Cat类型的同名函数
