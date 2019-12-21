package main

import (
	"fmt"
)

var (
	a int
)

func delcareZero(pointerVariable *int) {
	*pointerVariable = 0
	fmt.Println("& point variable", &pointerVariable)
	fmt.Println("simply print point variable", pointerVariable)
}

func main() {
	a := 5
	delcareZero(&a)
	fmt.Println("Simply print &a", &a)
	fmt.Println(a)

}
