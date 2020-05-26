package main

import "fmt"

/**
* variables in package
*/
var (
	aa = "hello"
	cc bool = false
	ss = "world"
)

func variableValueZero()  {
	var a int
	var s string

	fmt.Printf("%d, %q", a, s)
}

func variableInitial()  {
	a, b := 3, 233
	var s = "well done"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, s = 3, 233, true, "world"
	fmt.Println(a, b, c, s)
}

func variableTypeShorter()  {
	a, b, c, s := 3, 233, true, "world"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello")

	variableValueZero()

	variableInitial()

	variableTypeDeduction()

	variableTypeShorter()

	fmt.Println(aa, ss, cc)
}
