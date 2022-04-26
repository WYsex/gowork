package main

import (
	"errors"
	"fmt"
)

var (
	a             int    = 123
	b             int8   = 4
	s             string = "hello"
	c             rune   = 'A' //查看当前字母的ascii码
	t             bool   = true
	d, m, n       rune   = 'a', 'b', 'c' //多个变量在同一行
	f                    = 14            //省略变量类型
	q                    = int8(32)      //显式类型转型
	ErrShortWrite        = errors.New("short write")
)

func bar() {
	{ // 等价于第一个if的隐式代码块
		a := 1 // 变量a作用域始于此
		if false {
		} else {
			{ // 等价于第一个else if的隐式代码块
				b := 2 // 变量b的作用域始于此
				if false {

				} else {
					{ // 等价于第二个else if的隐式代码块
						c := 3 // 变量c作用域始于此
						if false {

						} else {
							println(a, b, c)
						}
						// 变量c的作用域终止于此
					}
				} // 变量b的作用域终止于此
			}
		} // 变量a作用域终止于此
	}
}
func main() {
	bar()
	fmt.Println(a, b, s, c, t)
	fmt.Println(d, m, n)
	fmt.Println(f)
	fmt.Println(q)
	i := 1 //短变量类型声明
	fmt.Println(i)
	h, k := 13, "你好" //多个短变量声明
	fmt.Println(h, k)
	fmt.Println(ErrShortWrite)

}
