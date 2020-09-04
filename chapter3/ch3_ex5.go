package main
import "fmt"

func main() {
	var val float64
	fmt.Println("input value in Fahrenheit")
	fmt.Scanf("%f", &val)
	val_conv := (val - 32.0)*5.0/9.0
	fmt.Println("value in Celsius: ", val_conv)
}
