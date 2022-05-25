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

type Custom struct {
	//name, pwd string
}

func (cus Custom) loging(name, pwd string) (fanhui string) {
	if name == "tom" && pwd == "123" {
		return "登录成功"
	} else {
		return "账号或密码错误"
	}
}
func main() {
	per1 := Custom{}
	b := per1.loging("tom", "1123")
	fmt.Println(b)

	per := Person2{
		name: "xiaohong",
	}
	per.sleep()
	per.eat()

}
