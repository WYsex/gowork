package main

import "fmt"

type Person3 struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person3, error) {
	if name == "" {
		return nil, fmt.Errorf("name not nil")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age not <=0")
	}
	return &Person3{name: name, age: age}, nil
}
func main() {
	spr, err := NewPerson("name", -1)
	if err == nil {
		fmt.Println("person:", spr)
	} else {
		fmt.Println(err)
	}

}
