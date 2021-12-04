package main

import (
	"belajar-go/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hello world!!")
	sentence := testAja()
	fmt.Println(sentence)

	resultadd := calculation.Add(26, 19)

	fmt.Println(resultadd)
}
