package main

import ("fmt"; "container/list")

func main() {
	var ls list.List

	ls.PushBack("hello")
	ls.PushBack("hello")
	ls.PushBack(23.89)
	for n := ls.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
}