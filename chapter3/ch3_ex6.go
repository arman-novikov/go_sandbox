package main
import "fmt"

func main() {
	var feet float64
	fmt.Println("input feets: ")
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Println("meters = ", meters)
}
