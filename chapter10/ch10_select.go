package main
import "fmt"
import "time"

func writer(ch chan<- string, timeout_s int, msg string) {
	for {
		ch <- msg
		time.Sleep(time.Second * time.Duration(timeout_s))
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go writer(c1, 3, "from 1")
	go writer(c2, 5, "from 2")

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			// default:
			// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

}