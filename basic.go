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

// 枚举类型
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)

	// 普通枚举类型
	fmt.Println(cpp, java, python, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	// 子增值枚举类型
	fmt.Println(b, kb, mb, gb, tb)
}
func main() {
	//fmt.Println("Hello world")
	//variable()
	//variableShorter()
	//variableInitialValue()
	enums()
}
