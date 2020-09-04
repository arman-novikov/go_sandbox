package main

import ("fmt"; "io/ioutil")

const FILE_NAME string = "test.txt"

func main() {
	bs, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		fmt.Println("can't ReadFile ", FILE_NAME)
		return
	}

	str := string(bs)
	fmt.Println(str)
}