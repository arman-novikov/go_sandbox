package main
import "fmt"

func printer(vals []int) {
	for _, v := range vals {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func add_10(vals []int) {
	for i, v := range vals {
		vals[i] = v+10
	}
}

func main() {
	vals := []int{1,2,3,4,5,6}
	simple(vals)
	printer(vals)
}
