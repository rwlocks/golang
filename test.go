package main
import "fmt"

func test(i int) int {
	return i+1
}

func main(){
	fmt.Println("result: ", test(12))
}
