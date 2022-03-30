package main

import "fmt"

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n",
		a, s)

}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func variableInitialValue() {
	var a = 3
	var s = "abc"
	fmt.Println(a, s)
}

func main() {
	//fmt.Println("Hello world")
	variable()
	variableShorter()
	variableInitialValue()
}
