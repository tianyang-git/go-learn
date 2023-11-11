package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		res := fmt.Sprintf("loop %d", i)
		fmt.Println(res)
		i = i + 1
	}

	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop for")
		break
	}

	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
