package main
import "fmt"

func test(i int) int {
	return i+1
}

type Mystr struct{
	b int
}

func main(){
	fmt.Println("result: ", test(12))
}
