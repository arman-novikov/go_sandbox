package main
import "fmt"

func main() {
	i := 199
	var ptr *int
	ptr = new(int)
	fmt.Println(*ptr)
	ptr = new(int)
	fmt.Println(*ptr)
	ptr = &i
	fmt.Println(*ptr)
}
