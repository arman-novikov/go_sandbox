package main
import "fmt"

func biggest_of(vals ... int) int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	res := vals[0]
	for i:= 1; i < len(vals); i++ {
		if vals[i] > res {
			res = vals[i]
		}
	}
	return res
}

func main() {
	fmt.Println(biggest_of(1,3,55,66,2))
	fmt.Println(biggest_of())
	fmt.Println(biggest_of(1))
}
