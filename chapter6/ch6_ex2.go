package main
import "fmt"

func half(v int) (int, bool) {
	return v / 2, v % 2 == 0
}

func main() {
	fmt.Println(half(10))
	fmt.Println(half(7789))
	fmt.Println(half(0))
}
