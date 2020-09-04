package main
import ("fmt"; "sort")

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	temp := ps[i]
	ps[i] = ps[j]
	ps[j] = temp
}

type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	temp := ps[i]
	ps[i] = ps[j]
	ps[j] = temp
}


func main() {
	people := []Person {
		{"Igor", 22},
		{"Petya", 12},
		{"Vasya", 72},
		{"Oleg", 19},
		{"Stas", 16},
	}

	sort.Sort(ByName(people))
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}