package a

import (
	"fmt"
)

func test() {
	a := -10
	fmt.Println(abs(a))
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
