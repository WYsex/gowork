package main

import "fmt"

func func1(a int, b int) (c int, d int) {
	n := a + 1
	g := b + 2
	return n + 1, g + 1

}

func main() {
	h, g := func1(1, 2)
	fmt.Println(h, g)
}
