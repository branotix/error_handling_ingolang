package main

import (
	"errors"
	"fmt"
)

func doubleeven(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("its number is odd")
	}
	return (num * 2), nil
}
func main() {
	reslut, err := doubleeven(3)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(reslut)
}
