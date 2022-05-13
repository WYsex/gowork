package main

import "fmt"

type Person2 struct {
	name string
}

func (per Person2) eat() {
	fmt.Printf("%v在吃饭", per.name)
}

func (per Person2) sleep() {
	fmt.Printf("%v睡着了", per.name)
}

func main() {
	per := Person2{
		name: "xiaohong",
	}
	per.sleep()
	per.eat()

}
