package main

import "fmt"

func sayhello(name string) {
	fmt.Printf("%s Hello\n", name)
}

func mg(name string, f func(a string)) {
	f(name)
}

func add(a, b int) int {
	return a + b
}

func jian(a, b int) int {
	return a - b
}

func jisuan(fuhao string) func(a, b int) int {
	switch fuhao {
	case "+":
		return add
	case "-":
		return jian
	default:
		return nil
	}

}

func main() {
	mg("wang", sayhello)
	ff := jisuan("+")
	r := ff(2, 1)
	fmt.Println(r)

}
