package main
import "fmt"

func main() {
	x := 5
	x += 1
	if x == 6 {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}
