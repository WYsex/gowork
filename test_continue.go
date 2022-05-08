package main

func test3() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
		to:
			if i == 1 && j == 1 {
				break to
			}
		}

	}
}

func main() {

}
