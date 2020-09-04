package main

import ("fmt"; "os")

const FILE_NAME string = "test.txt"

func main() {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println("can't Open ", FILE_NAME)
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("can't stat ", FILE_NAME)
		return
	}

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("can't Read ", FILE_NAME)
		return
	}

	str := string(bs)
	fmt.Println(str)
}