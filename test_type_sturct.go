package main

import "fmt"

type Person1 struct {
	id, num     int
	name, email string
}

func main() {
	var ceshi Person1
	ceshi.id = 1
	ceshi.num = 123
	ceshi.name = "zhangsan"
	ceshi.email = "zhangsan@gmail.com"
	fmt.Println(ceshi)
}
