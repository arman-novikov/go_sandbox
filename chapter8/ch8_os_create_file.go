package main

import ("fmt"; "os")

const FILE_NAME string = "test.md"
const BUFF string = "### it is ok"

func main() {
	file, err := os.Create(FILE_NAME)
	if (err != nil) {
		fmt.Println("can't Create ", FILE_NAME)
		return
	}
	defer file.Close()

	file.WriteString(BUFF)
}