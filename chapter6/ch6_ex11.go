package main
import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 10
	y := 20
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
