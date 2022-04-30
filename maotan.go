package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

func main() {
	table := "/Users/wang/gowork/普卡会员-新零售.xlsx"
	f, err := excelize.OpenFile(table)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(f)

	cols, err1 := f.GetCols("普")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	name1 := cols[1]
	name := name1[1:]
	//fmt.Println(name)
	name2 := name[1:]
	fmt.Println(name2)

	for _, v := range name2 {
		a := strings.TrimSpace(v)
		fmt.Println(a)
		g := f.NewSheet(v)
		//	f.Save()
		fmt.Println(g)
	}

	//f.SaveAs("/Users/wang/gowork/普卡会员-新零售.xlsx")

}
