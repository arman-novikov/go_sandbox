package main
import "fmt"

func sum(cont []int) int {
	res := 0
	for _, v := range cont {
		res += v
	}
	return res
}

func main() {
	vals := []int{1,2,3,4,5,6}
	fmt.Println(sum(vals))
}
