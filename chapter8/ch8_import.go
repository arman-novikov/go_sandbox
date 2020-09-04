package main

import "fmt"
import "mymath"

func main() {
	vals := []float64{1,2,3,4,5,6,7,8,9,88.567}
	fmt.Println(mymath.Average(vals))
	fmt.Println(mymath.Average([]float64{}))
}