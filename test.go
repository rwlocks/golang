package main
import "fmt"

func test(i int) int {
	return i+1
}

type Mystr struct{
	b int
	c string
}

func (s *Mystr) Test(a int, d string)() {
	s.b = a
	s.c = d
}

func main(){
	s := Mystr{100, "hello"}
	fmt.Println("result: ", s.b)
    s.Test(1000, "world")
	fmt.Println("result: ", s.c)
}
