package main
import "fmt"

func main() {
	x := make([]int,3,9)
	if len(x) == 3 {
		fmt.Println("len is ok")
	} else {
		fmt.Println("len is not ok")
	}
}
