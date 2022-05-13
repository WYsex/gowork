package main

import "fmt"

type Dog struct {
	name, color string
	age         int
}

type Person struct {
	dog  Dog
	name string
	age  int
}

//func main() {
//	jinmao := Dog{
//		name:  "heibei",
//		color: "white",
//		age:   2,
//	}
//	xiaohong := Person{
//		dog:  jinmao,
//		name: "xiaohong",
//		age:  20,
//	}
//	zong := xiaohong
//	fmt.Println(zong)
//}

func main() {
	var per Person
	per.dog.name = "heibei"
	per.dog.age = 2
	per.dog.color = "white"
	per.name = "xiaohong"
	per.age = 12
	fmt.Println(per)

}
