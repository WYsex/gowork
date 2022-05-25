package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showmsg(i int) {
	defer wp.Done()
	fmt.Printf("i的值%d\n", i)
}

func main() {
	for i := 0; i <= 10; i++ {
		go showmsg(i)
		wp.Add(1)
	}
	defer wp.Wait()
	fmt.Println("end...")

}
