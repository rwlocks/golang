package main

import (
	"fmt"
)


func test(i, j int) (x, y int) {
	x = j
	y = i
	return
}

func main() {
	var x, y = test(10, 20)
	fmt.Println(x, y)
}
