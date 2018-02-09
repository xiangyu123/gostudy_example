package main

import "fmt"

type TestInf interface {
	Echos(astr string)
}

type testStruct struct {
	name string
	age  int8
}

func (r *testStruct) Echos(astr string) {
	fmt.Printf("%s: %s\n", astr, r.name)
}

func main() {
	testdata := testStruct{
		"xuzhigui",
		18,
	}

	testdata.Echos("hello my name is")
}
