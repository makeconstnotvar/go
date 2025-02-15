package main

import (
	"fmt"
	"leet/binary"
)

func main() {
	//var result = twoSum([]int{2, 7, 11, 15}, 9)

	//fmt.Println("Индексы:", result)

	//Rnd()
	//inorderTraversal(getTestData2())
	var idx = binary.Search(binary.GetTestData())
	fmt.Println(idx)
}
