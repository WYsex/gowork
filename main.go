package main

import (
	"fmt"
	"github.com/google/uuid" //新增依赖
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hello word")
	logrus.Println("hellow ads")
	logrus.Println(uuid.NewString())
}
