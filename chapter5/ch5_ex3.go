package main
import "fmt"

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	expected := []string{"c", "d", "e"}
	
	if len(expected) != len(x[2:5]) {
		fmt.Println("not ok")
		return
	}
	for i, v := range x[2:5] {
		if expected[i] != v {
			fmt.Println("not ok")
			return
		}
	}
	fmt.Println("ok")
}
