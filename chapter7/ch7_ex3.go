package main
import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2.0)
}

func (r *Rectangle) area() float64 {
	return (r.y2-r.y1) * (r.x2-r.x1)
}

func (c *Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2.0 * ((r.y2-r.y1) + (r.x2-r.x1))
}

type Shape interface {
	area() float64
	perimeter() float64
}

func showArea(shapes ...Shape) {
	for i := 0; i < len(shapes); i++ {
		fmt.Println(shapes[i].area())
	}
}

func showPerimeter(shapes ...Shape) {
	for i := 0; i < len(shapes); i++ {
		fmt.Println(shapes[i].perimeter())
	}
}

func main() {
	c := Circle{10.0, 10.0, 8.0}
	r := Rectangle{1.0, 10.0, 1.0, 10.0}

	showArea(&c, &r)
	showPerimeter(&c, &r)
}