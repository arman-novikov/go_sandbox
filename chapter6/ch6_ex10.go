package main
import "fmt"
import "math"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)
	if math.Abs(x - 2.25) < 0.000001 {
		fmt.Println("all is ok")
	}
}
