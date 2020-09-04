package main
import "fmt"

type Person struct {
	Name string	
}

func (p *Person) Talk() {
	fmt.Println("Hi! My name is ", p.Name)
}

func main() {
	p := Person{Name: "Anton"}
	p.Talk()
}